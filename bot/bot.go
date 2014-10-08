package bot

import (
	"github.com/LapisBot/Lapislazuli/bot/irc"
	"github.com/LapisBot/Lapislazuli/bot/log"
	"github.com/LapisBot/Lapislazuli/config"
	"sync"
)

const Name = "Lapislazuli v0.1dev"

type Bot struct {
	Config *config.Config
	irc    map[*config.Server]*irc.Bot
}

func Create(config *config.Config) *Bot {
	return &Bot{
		Config: config,
	}
}

func (bot *Bot) Start() {
	log.Info("Starting Bot:", Name)

	bot.irc = make(map[*config.Server]*irc.Bot)
	for _, server := range bot.Config.Servers {
		ircbot := irc.Create(server)
		bot.irc[server] = ircbot
		go ircbot.Connect()
	}
}

func (bot *Bot) Stop() {
	var wg sync.WaitGroup
	wg.Add(len(bot.irc))

	for _, ircbot := range bot.irc {
		go func() {
			ircbot.Disconnect()
			wg.Done()
		}()
	}

	wg.Wait()
}
