package sink

import "github.com/danstiner/structlog"

func NewNoop() structlog.Logger {
	return structlog.Logger{
		Sink: Noop{},
	}
}

type Noop struct{}

func (s Noop) Log(event structlog.Event) {}
