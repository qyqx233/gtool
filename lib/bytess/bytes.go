package bytess

import (
	"reflect"
	"unsafe"
)

func AppendData(bs []byte, p uintptr, size int) []byte {
	x := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: p,
		Len:  size,
		Cap:  size,
	}))
	return append(bs, x...)
}
