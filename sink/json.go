package sink

import (
	"encoding/json"
	"io"
	"time"

	"github.com/danstiner/structlog"
)

func NewJson(writer io.Writer) structlog.Sink {
	return Json{json.NewEncoder(writer)}
}

type Json struct {
	encoder *json.Encoder
}

func (s Json) Log(event structlog.Event) {
	event.Fields["$level"] = event.Level.String()
	event.Fields["$message"] = event.Message
	event.Fields["$template"] = event.Template
	event.Fields["$timestamp"] = event.Timestamp.UTC().Format(time.RFC3339)
	err := s.encoder.Encode(event.Fields)
	if err != nil {
		panic(err)
	}
	if event.Level == structlog.PanicLevel {
		panic(event.Message)
	}
}
