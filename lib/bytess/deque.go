package bytess

import "unsafe"

type Deque struct {
	noCopy
	buf      []byte
	len, cap int
}

func NewDeque(bs []byte) *Deque {
	return &Deque{
		buf: bs,
		cap: cap(bs),
	}
}

func (b *Deque) Reset() {
	b.len = 0
}

func (b *Deque) Append(s string) {
	if b.len+len(s) > b.cap {
		buf := make([]byte, b.len*2+len(s))
		copy(buf, b.buf)
		copy(buf[b.len:], s)
		b.buf = buf
		b.cap = len(buf)
	} else {
		copy(b.buf[b.len:], s)
	}
	b.len += len(s)
}

func (b *Deque) Insert(s string) {
	if b.len+len(s) > b.cap {
		buf := make([]byte, b.len*2+len(s))
		copy(buf, b.buf)
		copy(buf[b.len:], s)
		b.buf = buf
		b.cap = len(buf)
	} else {
		copy(b.buf[len(s):], b.buf[:b.len])
		copy(b.buf, s)
	}
	b.len += len(s)
}

func (b *Deque) Len() int {
	return b.len
}

func (b *Deque) Cap() int {
	return b.cap
}

func (b *Deque) WriteBytes(s []byte) {
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

func (b *Deque) String() string {
	return *(*string)(unsafe.Pointer(&b.buf))
}
