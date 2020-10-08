package sql

import (
	"unsafe"
)

type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

type ByteBuffer struct {
	noCopy
	buf      []byte
	len, cap int
}

func NewByteBuffer(bs []byte) *ByteBuffer {
	return &ByteBuffer{
		buf: bs,
		cap: cap(bs),
	}
}

func (b *ByteBuffer) WriteString(s string) {
	if b.len+len(s) > b.cap {
		buf := make([]byte, b.len*2+len(s))
		copy(buf, b.buf)
		copy(buf[b.len:], s)
		b.buf = buf
		b.cap = len(buf)
	} else {
		copy(b.buf, s)
	}
	b.len += len(s)
}

func (b *ByteBuffer) Len() int {
	return b.len
}

func (b *ByteBuffer) Cap() int {
	return b.cap
}

func (b *ByteBuffer) WriteBytes(s []byte) {
	if b.len+len(s) > b.cap {
		buf := make([]byte, b.len*2+len(s))
		copy(buf, b.buf)
		copy(buf[b.len:], s)
		b.buf = buf
		b.cap = len(buf)
	} else {
		copy(b.buf, s)
	}
	b.len += len(s)
}

func (b *ByteBuffer) String() string {
	return *(*string)(unsafe.Pointer(&b.buf))
}
