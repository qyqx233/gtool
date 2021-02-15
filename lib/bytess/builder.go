package bytess

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

func (b *ByteBuffer) ensureCap(n int) {
	if b.len+n > b.cap {
		buf := make([]byte, b.len*2+n)
		copy(buf, b.buf)
		b.buf = buf
		b.cap = len(buf)
	}
}

func (b *ByteBuffer) WriteString(s string) {
	b.ensureCap(len(s))
	copy(b.buf[b.len:], s)
	b.len += len(s)
}

func (b *ByteBuffer) Len() int {
	return b.len
}

func (b *ByteBuffer) Cap() int {
	return b.cap
}

func (b *ByteBuffer) WriteBytes(s []byte) {
	b.ensureCap(len(s))
	copy(b.buf[b.len:], s)
	b.len += len(s)
}

func (b *ByteBuffer) String() string {
	buf := b.buf[:b.len]
	return *(*string)(unsafe.Pointer(&buf))
}

func (b *ByteBuffer) Bytes() []byte {
	return b.buf[:b.len]
}
