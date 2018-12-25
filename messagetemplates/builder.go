package messagetemplates

import (
	"strconv"
	"unsafe"
)

// Specialized string builder for maximum efficiency
type builder struct {
	buf []byte
}

func makeBuilder(capacity int) builder {
	return builder{
		buf: make([]byte, 0, capacity),
	}
}

func (b *builder) write(bytes []byte) {
	b.buf = append(b.buf, bytes...)
}

func (b *builder) writeInt(i int64) {
	b.buf = strconv.AppendInt(b.buf, i, 10)
}

func (b *builder) writeString(str string) {
	b.buf = append(b.buf, str...)
}

func (b *builder) string() string {
	return *(*string)(unsafe.Pointer(&b.buf))
}
