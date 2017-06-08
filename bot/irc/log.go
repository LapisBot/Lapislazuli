package irc

import "github.com/LapisBot/Lapislazuli/bot/log"

type ircLogger struct{}

func (l ircLogger) Debug(format string, args ...interface{}) {
	log.Debugf(format, args...)
}
func (l ircLogger) Info(format string, args ...interface{}) {
	log.Infof(format, args...)
}
func (l ircLogger) Warn(format string, args ...interface{}) {
	log.Warnf(format, args...)
}
func (l ircLogger) Error(format string, args ...interface{}) {
	log.Errorf(format, args...)
}
