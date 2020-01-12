package bytess

import (
	"testing"
)

type ByteBuffer1 struct {

	// B is a byte buffer to use in append-like workloads.
	// See example code for details.
	B []byte
}

// Len returns the size of the byte buffer.
func (b *ByteBuffer1) Len() int {
	return len(b.B)
}

func (b *ByteBuffer1) WriteString(s string) (int, error) {
	b.B = append(b.B, s...)
	return len(s), nil
}

func BenchmarkC1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bs := NewByteBuffer(make([]byte, 20))
		bs.WriteString("aaaabbbb")
	}
}

func BenchmarkC2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bs := NewBytesIter(make([]byte, 20))
		bs.WriteString("aaaabbbb")
		_ = bs.IterString()
	}
}

func Test2x(t *testing.T) {
	bi := NewBytesIter(make([]byte, 30))
	array := [30]byte{}
	var u uint64 = 21
	bi.WriteString("def")
	bi.WriteUint64(u)
	bi.WriteString("abc")
	t.Log(bi.IterString())
	t.Log(bi.IterUint64())
	t.Log(bi.IterString())
	bi.Dump(array[:])
	t.Log(array)
	as := NewBytesIter(make([]byte, 20))
	as.Load(array[:], bi.wo)
	t.Log(as.IterString())
	t.Log(as.wo)
	t.Log(as.Dump(array[:]))
	// bi.WriteString("asdfas")
	// t.Log(bi.Bytes())
	// t.Log(time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))
}

func Test3(t *testing.T) {
	var array []byte
	array = append(array, []byte("abc")...)
	bi := NewBytesIter(make([]byte, 10))
	bi.WriteString("abcd")
	// t.Log(bi.buf)
	sbi := (&BytesIter{}).SetBytes(array, 4)
	t.Log(sbi.IterString())
}
