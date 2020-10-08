package convert

import (
	"testing"
)

func Test(t *testing.T) {
	var u uint64 = 454
	t.Log(Uint642Bytes(u))
	var i int64 = 30
	t.Log(Int642Bytes(i))
	t.Log(Bytes2Uint64(Uint642Bytes(u)))
}
