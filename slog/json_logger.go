package slog

import (
	"io"

	"github.com/pkg/errors"
)

type jsonLogger struct {
	writer io.Writer
}

func (l jsonLogger) Debug(format string, values ...interface{}) {
}
func (l jsonLogger) Error(format string, values ...interface{}) {
}
func (l jsonLogger) Info(format string, values ...interface{}) {
}
func (l jsonLogger) Warn(format string, values ...interface{}) {
}
func (l jsonLogger) With(key, value interface{}) Logger {
	return l
}
func (l jsonLogger) WrapError(err error, format string, values ...interface{}) error {
	return errors.Wrap(err, format)
}

func NewJsonLogger(w io.Writer) Logger {
	return jsonLogger{
		writer: w,
	}
}
