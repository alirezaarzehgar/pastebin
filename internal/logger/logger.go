package logger

type Leveler uint

const (
	LevelError = iota
	LevelWarn
	LevelInfo
	LevelDebug
	LevelVerbose

	LevelUnknown
)

type Logger interface {
	Error(msg string, args ...any)
	Warn(msg string, args ...any)
	Info(msg string, args ...any)
	Debug(msg string, args ...any)
	Verbose(msg string, args ...any)
}

func Error(msg string, args ...any) {
	DefaultLogger.Error(msg, args...)
}

func Warn(msg string, args ...any) {
	DefaultLogger.Warn(msg, args...)
}

func Info(msg string, args ...any) {
	DefaultLogger.Info(msg, args...)
}

func Debug(msg string, args ...any) {
	DefaultLogger.Debug(msg, args...)
}

func Verbose(msg string, args ...any) {
	DefaultLogger.Verbose(msg, args...)
}
