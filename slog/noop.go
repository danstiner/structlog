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
func (l noopLogger) WrapError(err error, format string, values ...interface{}) error {
	return err
}

func NewNoopLogger() Logger {
	return noopLogger{}
}
