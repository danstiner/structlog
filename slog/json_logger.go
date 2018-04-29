package slog

import (
	"encoding/json"
	"io"

	"github.com/danstiner/go-structlog/messagetemplates"
)

type jsonLogger struct {
	writer io.Writer
	with   []KeyValue
}

func (l jsonLogger) Debug(format string, values ...interface{}) {
	l.log("DEBUG", format, values...)
}
func (l jsonLogger) Error(format string, values ...interface{}) {
	l.log("ERROR", format, values...)
}
func (l jsonLogger) Info(format string, values ...interface{}) {
	l.log("INFO", format, values...)
}
func (l jsonLogger) Warn(format string, values ...interface{}) {
	l.log("WARN", format, values...)
}
func (l jsonLogger) log(level string, format string, values ...interface{}) {
	msg, m, err := messagetemplates.Format(format, values...)
	if err != nil {
		panic(err)
	}

	for _, kv := range l.with {
		m[kv.key] = kv.value
	}

	m["@level"] = level
	m["@message"] = msg
	m["@template"] = format

	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	_, err = l.writer.Write(bytes)
	if err != nil {
		panic(err)
	}
}

func (l jsonLogger) With(key string, value interface{}) Logger {
	l.with = append(l.with, KV(key, value))
	return l
}

func NewJsonLogger(w io.Writer) Logger {
	return jsonLogger{
		writer: w,
	}
}
