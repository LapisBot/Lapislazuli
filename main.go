package main

import (
	"bufio"
	"fmt"
	"github.com/LapisBot/Lapislazuli/bot"
	"github.com/LapisBot/Lapislazuli/cli"
	"github.com/LapisBot/Lapislazuli/config"
	"github.com/ogier/pflag"
	"os"
	"os/signal"
	"path/filepath"
)

const (
	configName = "lapislazuli.json"
)

func Run(name string, args []string) int {
	flags := pflag.NewFlagSet(name, pflag.ContinueOnError)

	dir := flags.StringP("dir", "d", ".", "The folder to save all files in.")
	config := flags.StringP("config", "c", configName, "The configuration file used to configure the bot.")

	cli.FlagUsage(name, flags)

	if len(args) >= 1 && args[0] == "help" {
		flags.Usage()
		return 1
	}

	if flags.Parse(args) != nil {
		return 1
	}

	if *dir != "." && filepath.Dir(*config) == "." {
		*config = filepath.Join(*dir, *config)
	}

	// Load the configuration
	fmt.Println("Loading configuration from:", *config)
	conf := loadConfigFile(*config)
	if conf == nil {
		os.Exit(0)
	}

	// Launch the bot
	me := bot.Create(conf)

	// Make sure to shutdown gracefully if the program exists
	handleInterrupt(me.Stop)

	// Start the bot
	me.Start()
	defer me.Stop()

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = line[:len(line)-1]
		fmt.Println(line)
	}

	return 0
}

func main() {
	os.Exit(Run(filepath.Base(os.Args[0]), os.Args[1:]))
}

// This method should always complete successfully. If it doesn't, then something is really wrong and we
// shouldn't continue execution.
func assert(err error) {
	if err != nil {
		panic(err)
	}
}

// This method will exist the program if an error occurs and print the error message. Unlike assert() this
// is not that bad and won't panic the complete program.
func require(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
}

// Loads the configuration file from the specified file or creates a new empty one if it doesn't exist.
func loadConfigFile(path string) (conf *config.Config) {
	// Try opening the configuration file first so we can check if it exists
	file, err := os.Open(path)

	if os.IsNotExist(err) {
		// Create a new configuration file
		file, err = os.Create(path)
		require(err)
		defer file.Close()

		// Default values
		configFile := config.New()

		server := config.NewServer("example")
		server.Connection.Address = "irc.example.com:6697"
		server.Connection.SSL = true
		configFile.Servers["example"] = server
		configFile.Channels["example:#example"] = config.NewChannel()

		configFile.HTTP.Bind = ":8084"

		// Write the default configuration to the file
		require(config.Write(file, configFile))

		return nil
	} else {
		require(err)
		defer file.Close()

		// Read the configuration from the file
		conf, err = config.Parse(file)
		require(err)
	}

	return
}

func handleInterrupt(handler func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		handler()
		os.Exit(-1)
	}()
}
