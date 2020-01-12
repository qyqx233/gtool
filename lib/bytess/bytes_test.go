package bytess

import (
	"github.com/qyqx233/gtool/lib/assert"
	"testing"
)

type byteBuffer1 struct {

	// B is a byte buffer to use in append-like workloads.
	// See example code for details.
	B []byte
}

// Len returns the size of the byte buffer.
func (b *byteBuffer1) Len() int {
	return len(b.B)
}

func (b *byteBuffer1) WriteString(s string) (int, error) {
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

func BenchmarkC3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bs := &byteBuffer1{make([]byte, 0, 30)}
		bs.WriteString("aaaabbbb")
	}
}

func Test2x(t *testing.T) {
	array := [30]byte{}
	bi := NewBytesIter(array[:])
	var u uint64 = 21
	bi.WriteString("def")
	assert.AssertEqual(t, bi.wo, 4)
	bi.WriteUint64(u)
	assert.AssertEqual(t, bi.wo, 12)
	bi.WriteString("abc")
	bi.WriteByte(12)
	bi.WriteInt(100)
	assert.AssertTrue(t, bi.IterString() == "def")
	assert.AssertTrue(t, bi.IterUint64() == 21)
	assert.AssertTrue(t, bi.IterString() == "abc")
	assert.AssertTrue(t, bi.IterByte() == 12)
	assert.AssertTrue(t, bi.IterInt() == 100)
	bi1 := (&BytesIter{}).Loads(array[:], bi.wo)
	assert.AssertTrue(t, bi1.wo == bi.wo)
	assert.AssertTrue(t, bi1.IterString() == "def")
	bi1.WriteString("xxx")
	assert.AssertTrue(t, bi1.IterUint64() == 21)
	assert.AssertTrue(t, bi1.IterString() == "abc")
	// t.Log(bi1.Bytes())
	assert.AssertTrue(t, bi1.IterByte() == 12)
	assert.AssertTrue(t, bi1.IterInt() == 100)
	assert.AssertTrue(t, bi1.IterString() == "xxx")
}
