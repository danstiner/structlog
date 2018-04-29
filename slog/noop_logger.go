package slog

type noopLogger struct{}

func (l noopLogger) Debug(format string, values ...interface{}) {
}
func (l noopLogger) Error(format string, values ...interface{}) {
}
func (l noopLogger) Info(format string, values ...interface{}) {
}
func (l noopLogger) Warn(format string, values ...interface{}) {
}
func (l noopLogger) With(key, value interface{}) Logger {
	return l
}

func NewNoopLogger() Logger {
	return noopLogger{}
}
