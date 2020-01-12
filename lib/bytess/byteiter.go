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
	buf    uintptr
	wo, ro int
	cap    uintptr
}

func NewBytesIter(buf []byte) *BytesIter {
	bi := new(BytesIter)
	bi.buf = uintptr(unsafe.Pointer((&buf)))
	bi.cap = uintptr(len(buf))
	return bi
}

func (b *BytesIter) Reset() {
	b.wo = 0
	b.ro = 0
}

func (b *BytesIter) SetBytes(bs []byte, n int) *BytesIter {
	b.buf = uintptr(unsafe.Pointer(&bs))
	b.wo = n
	b.ro = 0
	return b
}

func (bi *BytesIter) Bytes() []byte {
	// return *(*[]byte)(unsafe.Pointer(&byteSliceStru{bi.buf, bi.cap, bi.cap}))
	return *(*[]byte)(unsafe.Pointer(bi.buf))
}

func (bi *BytesIter) WriteUint64(u uint64) int {
	// var l = int(unsafe.Sizeof(u))
	// sl := *(*[]byte)(unsafe.Pointer(&byteSliceStru{bi.buf, bi.cap, bi.cap}))
	// copy(sl[bi.wo:], convert.Uint642Bytes(u))
	// bi.wo += l
	buf := *(*[]byte)(unsafe.Pointer(bi.buf))
	copy(buf[bi.wo:], convert.Uint642Bytes(u))
	bi.wo += int(unsafe.Sizeof(u))
	return bi.wo
}

func (bi *BytesIter) WriteString(s string) int {
	// var l = len(s)
	// buf := *(*[]byte)(unsafe.Pointer(&byteSliceStru{bi.buf, bi.cap, bi.cap}))
	// buf[bi.wo] = byte(l)
	// copy(buf[bi.wo+1:], s)
	// bi.wo += (l + 1)
	var l = len(s)
	buf := *(*[]byte)(unsafe.Pointer(bi.buf))
	buf[bi.wo] = byte(l)
	copy(buf[bi.wo+1:], s)
	bi.wo += (l + 1)
	return bi.wo
}

func (bi *BytesIter) WriteBytes(s []byte) int {
	var l = len(s)
	// buf := *(*[]byte)(unsafe.Pointer(&byteSliceStru{bi.buf, bi.cap, bi.cap}))
	buf := *(*[]byte)(unsafe.Pointer(bi.buf))
	buf[bi.wo] = byte(l)
	copy(buf[bi.wo+1:], s)
	bi.wo += (l + 1)
	return bi.wo
}

func (bi *BytesIter) IterString() string {
	// buf := *(*[]byte)(unsafe.Pointer(&byteSliceStru{bi.buf, bi.cap, bi.cap}))
	buf := *(*[]byte)(unsafe.Pointer(bi.buf))
	if bi.wo <= bi.ro {
		return ""
	}
	var n = int(buf[bi.ro])
	s := *(*string)(unsafe.Pointer(
		&stringStru{uintptr(unsafe.Pointer(&buf[bi.ro+1])), n}))
	bi.ro += (n + 1)
	return s
}

func (bi *BytesIter) IterStringSafe() string {
	// buf := *(*[]byte)(unsafe.Pointer(&byteSliceStru{bi.buf, bi.cap, bi.cap}))
	buf := *(*[]byte)(unsafe.Pointer(bi.buf))
	if bi.wo <= bi.ro {
		return ""
	}
	var n = int(buf[bi.ro])
	s := string(buf[bi.ro+1 : n])
	bi.ro += (n + 1)
	return s
}

func (bi *BytesIter) IterUint64() uint64 {
	// buf := *(*[]byte)(unsafe.Pointer(&byteSliceStru{bi.buf, bi.cap, bi.cap}))
	// u := *(*uint64)(unsafe.Pointer(&buf[bi.ro]))
	// bi.ro += int(unsafe.Sizeof(uint64(0)))
	// return u
	buf := *(*[]byte)(unsafe.Pointer(bi.buf))
	u := *(*uint64)(unsafe.Pointer(&buf[bi.ro]))
	bi.ro += int(unsafe.Sizeof(uint64(0)))
	return u
}

func (bi *BytesIter) Dump(bs []byte) int {
	buf := *(*[]byte)(unsafe.Pointer(bi.buf))
	copy(bs, buf)
	return bi.wo
}

func (bi *BytesIter) Loads(bs []byte, n int) *BytesIter {
	// buf := *(*[]byte)(unsafe.Pointer(&bs))
	// copy(buf, bs)
	bi.buf = uintptr(unsafe.Pointer(&bs))
	bi.ro = 0
	bi.wo = n
	return bi
}
