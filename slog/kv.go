package slog

type KeyValue struct {
	key   string
	value interface{}
}

func KV(key string, value interface{}) KeyValue {
	return KeyValue{
		key:   key,
		value: value,
	}
}
