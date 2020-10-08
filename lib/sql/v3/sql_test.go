package sql

import (
	"gtool/lib/bytess"
	"strings"
	"testing"
)

type ByteSlice []byte

func (b *ByteSlice) WriteString(s string) {
	// *b = append(*b, s...)
	copy(*b, s)
}

func fy(bs *ByteSlice) {
	bs.WriteString("abc")
}

var long = strings.Repeat("a", 101)

func BenchmarkC2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := [100]byte{}
		bb := NewByteBuffer(buf[:])
		bb.WriteString("abcdef")
		bb.WriteString(long)
		// fy(&bs)
		// t.Log(buf)
		// bs.foo("a")
		// bs = bs.add(long)
	}
}

func BenchmarkC1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bs := bytess.NewByteBuffer(make([]byte, 10))
		bs.WriteString("aaaabbbb")
		// bb.WriteString(long)
		// fy(&bs)
		// t.Log(buf)
		// bs.foo("a")
		// bs = bs.add(long)
	}
}

func Test8(t *testing.T) {
	// buf := [10]byte{}

	bs := bytess.NewByteBuffer(make([]byte, 10))
	bs.WriteString("aaaabbbb")
	bs.WriteString(long)
	t.Log(bs.String())
	t.Log(bs.Cap(), bs.Len())
}
