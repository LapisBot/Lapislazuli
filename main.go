package main

import (
	"flag"
	"os"
	"fmt"
	"path/filepath"
	"github.com/LapisBot/Lapislazuli/bot"
	"github.com/LapisBot/Lapislazuli/config"
)

const Name = "Lapislazuli v0.1dev"

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Available flags:")

		// Print defaults
		flag.CommandLine.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(os.Stderr, "  -%s | %s - (%s)\n", f.Name, f.Usage, f.DefValue)
		})
	}

	// The folder to save the Bot files in
	dir, err := os.Getwd(); assert(err)
	flag.StringVar(&dir, "dir", dir, "Set the folder for all the Bot files")

	configFile := "config.json"
	flag.StringVar(&configFile, "config", configFile, "The file to load the configuration from")

	// Parse the flags given when running the application
	flag.Parse()

	// We have the parsed flags available now, load the configuration
	configFile = filepath.Join(dir, configFile)
	fmt.Println("Loading configuration from", configFile)
	conf := loadConfigFile(configFile)

	// Launch the bot
	fmt.Println("Starting", Name, "Bot...")
	bot.Create(conf).Start()
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
func require(err error) bool {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
	return true
}

// Loads the configuration file from the specified file or creates a new empty one if it doesn't exist.
func loadConfigFile(path string) (conf *config.Config) {
	// Try opening the configuration file first so we can check if it exists
	file, err := os.Open(path)
	if os.IsNotExist(err) {
		// Create a new configuration file
		file, err = os.Create(path); require(err)
		defer file.Close()

		// Default values
		conf = config.New()
		conf.Servers["irc.example.com:+6697"] = config.NewServer()

		// Write the default configuration to the file
		require(config.Write(file, conf))
	} else {
		require(err)
		defer file.Close()

		// Read the configuration from the file
		conf, err = config.Read(file); require(err)
	}

	return
}
