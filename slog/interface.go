package slog

type Logger interface {
	Log(format string, v ...interface{})
}

type LeveledLogger interface {
	Logger
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Error(format string, v ...interface{})
	WrapError(err error, format string, v ...interface{}) error
}
