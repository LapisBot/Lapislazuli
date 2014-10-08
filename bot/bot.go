package bot

import (
	"github.com/LapisBot/Lapislazuli/bot/irc"
	"github.com/LapisBot/Lapislazuli/bot/log"
	"github.com/LapisBot/Lapislazuli/config"
)

const Name = "Lapislazuli v0.1dev"

type Bot struct {
	Config *config.Config
	irc    *irc.IRCBot
}

func Create(config *config.Config) *Bot {
	return &Bot{
		Config: config,
	}
}

func (bot *Bot) Start() {
	log.Info("Starting Bot:", Name)

	for _, server := range bot.Config.Servers {
		bot.irc = irc.Create(server)
		bot.irc.Connect()
	}
}

func (bot *Bot) Stop() {
	bot.irc.Disconnect()
}
