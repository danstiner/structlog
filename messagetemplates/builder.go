package messagetemplates

import (
	"strconv"
	"unsafe"
)

// String builder specialized for this library. Similar to `strings.Builder`
// but allows specifying a starting capacity and other performance improvements.
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
