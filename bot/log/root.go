package log

var logger Logger = &goLogger{}

func Get() Logger {
	return logger
}

func Log(level Level, v ...interface{}) {
	logger.Log(level, v...)
}
func Logf(level Level, format string, args ...interface{}) {
	logger.Logf(level, format, args...)
}

func Trace(v ...interface{}) {
	Log(TraceLevel, v...)
}
func Tracef(format string, args ...interface{}) {
	Logf(TraceLevel, format, args...)
}

func Debug(v ...interface{}) {
	Log(DebugLevel, v...)
}
func Debugf(format string, args ...interface{}) {
	Logf(DebugLevel, format, args...)
}

func Info(v ...interface{}) {
	Log(InfoLevel, v...)
}
func Infof(format string, args ...interface{}) {
	Logf(InfoLevel, format, args...)
}

func Warn(v ...interface{}) {
	Log(WarnLevel, v...)
}
func Warnf(format string, args ...interface{}) {
	Logf(WarnLevel, format, args...)
}

func Error(v ...interface{}) {
	Log(ErrorLevel, v...)
}
func Errorf(format string, args ...interface{}) {
	Logf(ErrorLevel, format, args...)
}
