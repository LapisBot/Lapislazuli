package irc

import (
	"crypto/tls"
	"github.com/LapisBot/Lapislazuli/bot/log"
	"github.com/LapisBot/Lapislazuli/config"
	"github.com/fluffle/goirc/client"
	"github.com/fluffle/goirc/logging"
)

type IRCBot struct {
	Conn *client.Conn
	quit chan struct{}
}

func init() {
	logging.SetLogger(&ircLogger{})
}

func Create(server *config.Server) (bot *IRCBot) {
	nick := server.Login.User
	if nick == "" {
		nick = "Lapislazuli"
		log.Errorf("No nick specified, using %s as default", nick)
	}

	conf := client.NewConfig(nick, server.Login.Ident, server.Login.Name)
	conf.PingFreq = 0
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
	bot = &IRCBot{Conn: conn}
	conn.HandleFunc("connected", bot.connected)
	conn.HandleFunc("disconnected", bot.disconnected)
	return
}

func (irc *IRCBot) Connect() {
	irc.quit = make(chan struct{})
	irc.Conn.Connect()
	<-irc.quit
}

func (irc *IRCBot) Disconnect() {
	irc.Conn.Quit()
	<-irc.quit
}

func (irc *IRCBot) connected(conn *client.Conn, line *client.Line) {
	conn.Join("#lapislazuli") // TODO
}

func (irc *IRCBot) disconnected(conn *client.Conn, line *client.Line) {
	close(irc.quit)
}
