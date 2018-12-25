package slog

func Noop() Logger {
	return Logger{
		sink: NoopSink{},
	}
}

type NoopSink struct{}

func (s NoopSink) Log(event Event) {}
