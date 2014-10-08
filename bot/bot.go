package bot

import (
	"github.com/LapisBot/Lapislazuli/bot/irc"
	"github.com/LapisBot/Lapislazuli/bot/log"
	"github.com/LapisBot/Lapislazuli/config"
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
	for _, ircbot := range bot.irc {
		ircbot.Disconnect()
	}
}
