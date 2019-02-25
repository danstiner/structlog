package sink

import (
	"encoding/json"
	"fmt"
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
	m := make(map[string]interface{})

	for _, datum := range event.Data {
		m[fmt.Sprintf("%v", datum.Key)] = datum.Value
	}

	m["$level"] = event.Level.String()
	m["$message"] = event.Message
	m["$template"] = event.Template
	m["$timestamp"] = event.Timestamp.UTC().Format(time.RFC3339)

	err := s.encoder.Encode(m)
	if err != nil {
		panic(err)
	}
}
