package log

type Logger interface {
	Name() string
	Level() Level

	Log(level Level, v ...interface{})
	Logf(level Level, format string, args ...interface{})
}
