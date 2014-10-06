package log

type Logger interface {
	Name() string
	Level() Level

	Log(level Level, message string)
	Logf(level Level, format string, args ...interface{})
}
