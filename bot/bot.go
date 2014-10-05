package bot

import "github.com/LapisBot/Lapislazuli/config"

type Bot struct {
	Config *config.Config
}

func Create(config *config.Config) *Bot {
	return &Bot {
		config,
	}
}

func (bot *Bot) Start() {

}
