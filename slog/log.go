package slog

import (
	"fmt"

	structlog "github.com/danstiner/go-structlog"
	"github.com/danstiner/go-structlog/messagetemplates"
)

type Level int

const (
	TraceLevel Level = iota
	InfoLevel
	ErrorLevel
	PanicLevel
)

func (l Level) String() string {
	switch l {
	case TraceLevel:
		return "TRACE"
	case InfoLevel:
		return "INFO"
	case ErrorLevel:
		return "ERROR"
	case PanicLevel:
		return "PANIC"
	default:
		panic(fmt.Sprintf("Unknown log level: %d", l))
	}
}

type Event struct {
	data     []structlog.KV
	level    Level
	message  string
	template string
}

type LogSink interface {
	Log(event Event)
}

type Logger struct {
	context []structlog.KV
	level   Level
	sink    LogSink
}

func (l Logger) Level(level Level) Logger {
	l.level = level
	return l
}

func (l Logger) With(key string, value interface{}) Logger {
	l.context = append(l.context, structlog.KV{key, value})
	return l
}

func (l *Logger) Trace(template string, values ...interface{}) {
	l.sink.Log(l.event(TraceLevel, template, values...))
}

func (l *Logger) Info(template string, values ...interface{}) {
	l.sink.Log(l.event(InfoLevel, template, values...))
}

func (l *Logger) Error(template string, values ...interface{}) {
	l.sink.Log(l.event(ErrorLevel, template, values...))
}

func (l *Logger) Panic(template string, values ...interface{}) {
	event := l.event(PanicLevel, template, values...)
	l.sink.Log(event)
	panic(event)
}

func (l *Logger) event(level Level, template string, values ...interface{}) Event {
	data := l.context
	message, kv, err := messagetemplates.Format(template, values...)
	if err != nil {
		panic(err)
	}

	data = append(data, kv...)

	return Event{
		data,
		level,
		message,
		template,
	}
}
