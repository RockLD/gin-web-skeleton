package logger

var (
	DefaultLogger Logger = NewHelper(NewLogger())
)

type Logger interface {
	Init(options ...Option) error

	Options() Options

	Fields(fields map[string]interface{}) Logger

	Log(level Level, v ...interface{})

	Logf(level Level, format string, v ...interface{})

	String() string
}

func Init(opts ...Option) error {
	return DefaultLogger.Init(opts...)
}

func Fields(fields map[string]interface{}) Logger {
	return DefaultLogger.Fields(fields)
}

func Log(level Level, v ...interface{}) {
	DefaultLogger.Log(level, v...)
}

func Logf(level Level, format string, v ...interface{}) {
	DefaultLogger.Logf(level, format, v...)
}

func String() string {
	return DefaultLogger.String()
}
