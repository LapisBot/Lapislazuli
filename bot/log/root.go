package log

var logger Logger = &goLogger{}

func Get() Logger {
	return logger
}

func Log(level Level, message string) {
	logger.Log(level, message)
}

func Logf(level Level, format string, args ...interface{}) {
	logger.Logf(level, format, args...)
}

func Trace(message string) {
	Log(TraceLevel, message)
}

func Tracef(format string, args ...interface{}) {
	Logf(TraceLevel, format, args...)
}

func Debug(message string) {
	Log(DebugLevel, message)
}

func Debugf(format string, args ...interface{}) {
	Logf(DebugLevel, format, args...)
}

func Info(message string) {
	Log(InfoLevel, message)
}

func Infof(format string, args ...interface{}) {
	Logf(InfoLevel, format, args...)
}

func Warn(message string) {
	Log(WarnLevel, message)
}

func Warnf(format string, args ...interface{}) {
	Logf(WarnLevel, format, args...)
}

func Error(message string) {
	Log(ErrorLevel, message)
}

func Errorf(format string, args ...interface{}) {
	Logf(ErrorLevel, format, args...)
}
