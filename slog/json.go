package slog

import (
	"encoding/json"
	"fmt"
	"io"
)

func Json(writer io.Writer) Logger {
	return Logger{
		sink: JsonSink{writer},
	}
}

type JsonSink struct {
	Writer io.Writer
}

func (s JsonSink) Log(event Event) {
	m := make(map[string]interface{})

	for _, datum := range event.data {
		m[fmt.Sprintf("%v", datum.Key)] = datum.Value
	}

	m["@level"] = event.level.String()
	m["@message"] = event.message
	m["@template"] = event.template

	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	_, err = s.Writer.Write(bytes)
	if err != nil {
		panic(err)
	}
}
