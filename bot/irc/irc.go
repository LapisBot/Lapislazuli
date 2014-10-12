package irc

import (
	"crypto/tls"
	"github.com/LapisBot/Lapislazuli/bot/log"
	"github.com/LapisBot/Lapislazuli/config"
	"github.com/fluffle/goirc/client"
	"github.com/fluffle/goirc/logging"
)

type Bot struct {
	Config *config.Server
	conn *client.Conn
	quit chan struct{}
}

func init() {
	logging.SetLogger(&ircLogger{})
}

func Create(server *config.Server) (bot *Bot) {
	nick := server.Login.User
	if nick == "" {
		nick = "Lapislazuli"
		log.Errorf("No nick specified, using %s as default", nick)
	}

	conf := client.NewConfig(nick, server.Login.Ident, server.Login.Name)
	conf.Version = server.Messages.Version
	conf.QuitMessage = server.Messages.Quit

	conf.Server = server.Connection.Address
	conf.Pass = server.Login.Password
	conf.SSL = server.Connection.SSL
	if !server.Connection.Certificate {
		conf.SSLConfig = &tls.Config{}
		conf.SSLConfig.InsecureSkipVerify = true
	}

	conn := client.Client(conf)
	bot = &Bot{
		Config: server,
		conn: conn,
	}

	conn.HandleFunc("connected", bot.connected)
	conn.HandleFunc("disconnected", bot.disconnected)
	return
}

func (irc *Bot) Connect() {
	irc.quit = make(chan struct{})
	irc.conn.Connect()
	<-irc.quit
}

func (irc *Bot) Disconnect() {
	irc.conn.Quit()
	<-irc.quit
}

func (irc *Bot) connected(conn *client.Conn, _ *client.Line) {
	for channel, _ := range irc.Config.Channels {
		conn.Join(channel)
	}
}

func (irc *Bot) disconnected(_ *client.Conn, _ *client.Line) {
	close(irc.quit)
}
