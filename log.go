package structlog

import (
	"fmt"
	"time"

	"github.com/danstiner/structlog/messagetemplates"
)

func New(sink Sink) Logger {
	return Logger{
		Sink: sink,
	}
}

type Fields map[string]interface{}

type Logger struct {
	Fields Fields
	Sink   Sink
}

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

type Sink interface {
	Log(event Event)
}

type Event struct {
	Fields    Fields
	Level     Level
	Message   string
	Template  string
	Timestamp time.Time
}

const ErrorKey = "error"

func (l *Logger) With(key string, value interface{}) Logger {
	newFields := make(Fields, len(l.Fields)+1)
	for k, v := range l.Fields {
		newFields[k] = v
	}
	newFields[key] = value
	return Logger{newFields, l.Sink}
}

func (l *Logger) WithFields(fields Fields) Logger {
	newFields := make(Fields, len(l.Fields)+len(fields))
	for k, v := range l.Fields {
		newFields[k] = v
	}
	for k, v := range fields {
		newFields[k] = v
	}
	return Logger{newFields, l.Sink}
}

func (l *Logger) WithError(err error) Logger {
	return l.With(ErrorKey, err)
}

func (l *Logger) Trace(template string, values ...interface{}) {
	l.Sink.Log(l.event(TraceLevel, template, values...))
}

func (l *Logger) Info(template string, values ...interface{}) {
	l.Sink.Log(l.event(InfoLevel, template, values...))
}

func (l *Logger) Error(template string, values ...interface{}) {
	l.Sink.Log(l.event(ErrorLevel, template, values...))
}

func (l *Logger) Panic(template string, values ...interface{}) {
	event := l.event(PanicLevel, template, values...)
	l.Sink.Log(event)
	panic(event)
}

func (l *Logger) event(level Level, template string, values ...interface{}) Event {
	timestamp := time.Now()
	message, kvs, err := messagetemplates.Format(template, values...)
	if err != nil {
		panic(err)
	}
	fields := make(Fields, len(l.Fields)+len(kvs))
	for k, v := range l.Fields {
		fields[k] = v
	}
	for _, kv := range kvs {
		fields[kv.Key] = kv.Value
	}
	return Event{
		fields,
		level,
		message,
		template,
		timestamp,
	}
}
