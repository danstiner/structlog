package sink

import (
	"fmt"
	"time"

	"github.com/danstiner/structlog"
	"github.com/sirupsen/logrus"
)

type Logrus interface {
	WithFields(fields logrus.Fields) *logrus.Entry
	WithError(err error) *logrus.Entry
	WithTime(t time.Time) *logrus.Entry

	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

type LogrusSink struct {
	log Logrus
}

func NewLogrus(log Logrus) structlog.Sink {
	return LogrusSink{log}
}

func (s LogrusSink) Log(event structlog.Event) {
	var fields map[string]interface{} = event.Fields
	fields["$template"] = event.Template

	entry := s.log.WithFields(fields).WithTime(event.Timestamp)
	if value, ok := fields[structlog.ErrorKey]; ok {
		if err, ok := value.(error); ok {
			entry = entry.WithError(err)
		}
	}

	switch event.Level {
	case structlog.TraceLevel:
		entry.Trace(event.Message)
	case structlog.InfoLevel:
		entry.Info(event.Message)
	case structlog.ErrorLevel:
		entry.Error(event.Message)
	case structlog.PanicLevel:
		entry.Panic(event.Message)
	default:
		panic(fmt.Sprintf("Unknown log level: %d", event.Level))
	}
}
