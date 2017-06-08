package log

type Level struct {
	Prefix     string
	importance int8
}

var (
	AllLevels  = Level{importance: -128}
	TraceLevel = Level{"[TRACE]", -127}
	DebugLevel = Level{"[DEBUG]", -64}
	InfoLevel  = Level{"[INFO]", 0}
	WarnLevel  = Level{"[WARN]", 64}
	ErrorLevel = Level{"[ERROR]", 127}
)

func (level Level) IsLoggable(other Level) bool {
	return level.importance <= other.importance
}
