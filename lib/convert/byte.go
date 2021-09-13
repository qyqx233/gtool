package convert

import (
	"unsafe"
)

func Bytes2String(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

func String2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	b := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&b))
}

var isLittle bool = true

func init() {
	var u uint64 = 1
	if Uint642Bytes(u)[0] == 0 {
		isLittle = false
	}
}

func Bytes2Uint64(bs []byte) uint64 {
	if isLittle {
		return uint64(bs[0]) + (uint64(bs[1]) << 8) + (uint64(bs[2]) << 16) +
			(uint64(bs[3]) << 24) + (uint64(bs[4]) << 32) + (uint64(bs[5]) << 40) +
			(uint64(bs[6])<<48 + (uint64(bs[7]) << 56))
	} else {
		return uint64(bs[7]) + (uint64(bs[6]) << 8) + (uint64(bs[5]) << 16) +
			(uint64(bs[4]) << 24) + (uint64(bs[3]) << 32) + (uint64(bs[2]) << 40) +
			(uint64(bs[1])<<48 + (uint64(bs[0]) << 56))
	}
}

type byteSliceStru struct {
	buf      uintptr
	len, cap uintptr
}

func Uint642Bytes(u uint64) []byte {
	bs := byteSliceStru{}
	bs.buf = uintptr(unsafe.Pointer(&u))
	bs.len = 8
	bs.cap = 8
	return *(*[]byte)(unsafe.Pointer(&bs))
}

func Int642Bytes(u int64) []byte {
	bs := byteSliceStru{}
	bs.buf = uintptr(unsafe.Pointer(&u))
	bs.len = unsafe.Sizeof(u)
	bs.cap = unsafe.Sizeof(u)
	return *(*[]byte)(unsafe.Pointer(&bs))
}

func Int2Bytes(u int) []byte {
	bs := byteSliceStru{}
	bs.buf = uintptr(unsafe.Pointer(&u))
	bs.len = unsafe.Sizeof(u)
	bs.cap = unsafe.Sizeof(u)
	return *(*[]byte)(unsafe.Pointer(&bs))
}
