package slog

type Logger interface {
	Debug(format string, values ...interface{})
	Error(format string, values ...interface{})
	Info(format string, values ...interface{})
	Warn(format string, values ...interface{})
	With(key, value interface{}) Logger
	WrapError(err error, format string, values ...interface{}) error
}
