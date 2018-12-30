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

type Logger struct {
	Context []messagetemplates.KV
	Sink    Sink
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
	Data      []messagetemplates.KV
	Level     Level
	Message   string
	Template  string
	Timestamp time.Time
}

func (l Logger) With(key string, value interface{}) Logger {
	l.Context = append(l.Context, messagetemplates.KV{Key: key, Value: value})
	return l
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
	data := l.Context
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
		timestamp,
	}
}
