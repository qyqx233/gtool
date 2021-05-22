package bytess

import (
	"unsafe"

	"github.com/qyqx233/gtool/lib/convert"
)

type byteSliceStru struct {
	buf      uintptr
	len, cap uintptr
}

type stringStru struct {
	buf uintptr
	len int
}

type BytesIter struct {
	buf    []byte
	wo, ro int
	cap    uintptr
}

func NewBytesIter(buf []byte) *BytesIter {
	bi := new(BytesIter)
	bi.cap = uintptr(len(buf))
	return bi
}

func (b *BytesIter) Reset() {
	b.wo = 0
	b.ro = 0
}

func (bi *BytesIter) Bytes() []byte {
	return bi.buf[:bi.wo]
}

func (bi *BytesIter) BytesCopy() []byte {
	buf1 := make([]byte, bi.wo)
	copy(buf1, bi.buf)
	return buf1
}

func (bi *BytesIter) WriteUint64(u uint64) int {
	copy(bi.buf[bi.wo:], convert.Uint642Bytes(u))
	bi.wo += int(unsafe.Sizeof(u))
	return bi.wo
}

func (bi *BytesIter) WriteInt(u int) int {
	copy(bi.buf[bi.wo:], convert.Int2Bytes(u))
	bi.wo += int(unsafe.Sizeof(u))
	return bi.wo
}

func (bi *BytesIter) WriteInt64(u int64) int {
	copy(bi.buf[bi.wo:], convert.Int642Bytes(u))
	bi.wo += int(unsafe.Sizeof(u))
	return bi.wo
}

func (bi *BytesIter) WriteString(s string) int {
	var l = len(s)
	bi.buf[bi.wo] = byte(l)
	copy(bi.buf[bi.wo+1:], s)
	bi.wo += (l + 1)
	return bi.wo
}

func (bi *BytesIter) WriteBytes(s []byte) int {
	var l = len(s)
	// buf := *(*[]byte)(unsafe.Pointer(&byteSliceStru{bi.buf, bi.cap, bi.cap}))
	bi.buf[bi.wo] = byte(l)
	copy(bi.buf[bi.wo+1:], s)
	bi.wo += (l + 1)
	return bi.wo
}

func (bi *BytesIter) WriteByte(b byte) int {
	bi.buf[bi.wo] = b
	bi.wo += 1
	return bi.wo
}

func (bi *BytesIter) Write([]byte) (n int, err error) {
	// buf := *(*[]byte)(unsafe.Pointer(bi.buf))
	// buf[bi.wo] = b
	// bi.wo += 1
	return bi.wo, nil
}

func (bi *BytesIter) IterString() string {
	// buf := *(*[]byte)(unsafe.Pointer(&byteSliceStru{bi.buf, bi.cap, bi.cap}))
	if bi.wo <= bi.ro {
		return ""
	}
	var n = int(bi.buf[bi.ro])
	s := *(*string)(unsafe.Pointer(
		&stringStru{uintptr(unsafe.Pointer(&bi.buf[bi.ro+1])), n}))
	bi.ro += (n + 1)
	return s
}

func (bi *BytesIter) IterStringSafe() string {
	// buf := *(*[]byte)(unsafe.Pointer(&byteSliceStru{bi.buf, bi.cap, bi.cap}))
	if bi.wo <= bi.ro {
		return ""
	}
	var n = int(bi.buf[bi.ro])
	s := string(bi.buf[bi.ro+1 : n])
	bi.ro += (n + 1)
	return s
}

func (bi *BytesIter) IterUint64() uint64 {
	u := *(*uint64)(unsafe.Pointer(&bi.buf[bi.ro]))
	bi.ro += int(unsafe.Sizeof(uint64(0)))
	return u
}

func (bi *BytesIter) IterInt64() int64 {
	u := *(*int64)(unsafe.Pointer(&bi.buf[bi.ro]))
	bi.ro += int(unsafe.Sizeof(int64(0)))
	return u
}

func (bi *BytesIter) IterInt() int {
	u := *(*int)(unsafe.Pointer(&bi.buf[bi.ro]))
	bi.ro += int(unsafe.Sizeof(int(0)))
	return u
}

func (bi *BytesIter) IterByte() byte {
	// buf := *(*[]byte)(unsafe.Pointer(&byteSliceStru{bi.buf, bi.cap, bi.cap}))
	// u := *(*uint64)(unsafe.Pointer(&buf[bi.ro]))
	// bi.ro += int(unsafe.Sizeof(uint64(0)))
	// return u
	b := *(*byte)(unsafe.Pointer(&bi.buf[bi.ro]))
	bi.ro += int(unsafe.Sizeof(byte(0)))
	return b
}

func (bi *BytesIter) Dump(bs []byte) int {
	copy(bs, bi.buf)
	return bi.wo
}

func (bi *BytesIter) Loads(bs []byte, n int) *BytesIter {
	// buf := *(*[]byte)(unsafe.Pointer(&bs))
	// copy(buf, bs)
	bi.buf = bs[:n]
	bi.ro = 0
	bi.wo = n
	return bi
}
