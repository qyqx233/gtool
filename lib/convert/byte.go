package convert

import (
	"unsafe"
)

func Bytes2String(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
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
