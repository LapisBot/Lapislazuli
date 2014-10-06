package log

func LTrace(logger Logger, message string) {
	logger.Log(TraceLevel, message)
}

func LTracef(logger Logger, format string, args ...interface{}) {
	logger.Logf(TraceLevel, format, args...)
}

func LDebug(logger Logger, message string) {
	logger.Log(DebugLevel, message)
}

func LDebugf(logger Logger, format string, args ...interface{}) {
	logger.Logf(DebugLevel, format, args...)
}

func LInfo(logger Logger, message string) {
	logger.Log(InfoLevel, message)
}

func LInfof(logger Logger, format string, args ...interface{}) {
	logger.Logf(InfoLevel, format, args...)
}

func LWarn(logger Logger, message string) {
	logger.Log(WarnLevel, message)
}

func LWarnf(logger Logger, format string, args ...interface{}) {
	logger.Logf(WarnLevel, format, args...)
}

func LError(logger Logger, message string) {
	logger.Log(ErrorLevel, message)
}

func LErrorf(logger Logger, format string, args ...interface{}) {
	logger.Logf(ErrorLevel, format, args...)
}
