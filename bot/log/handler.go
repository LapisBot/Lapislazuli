package log

import golog "log"

type goLogger struct{}

func (logger *goLogger) Name() string {
	return "Logger"
}

func (logger *goLogger) Level() Level {
	return AllLevels
}

func (logger *goLogger) Log(level Level, v ...interface{}) {
	if logger.Level().IsLoggable(level) {
		golog.Println(append([]interface{}{level.Prefix}, v...)...)
	}
}

func (logger *goLogger) Logf(level Level, format string, args ...interface{}) {
	if logger.Level().IsLoggable(level) {
		golog.Printf(level.Prefix+" "+format+"\n", args...)
	}
}
