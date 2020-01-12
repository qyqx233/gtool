package bytess

import (
	"testing"

	"github.com/qyqx233/gtool/lib/assert"
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
		_ = bs.String()
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

func Test2y(t *testing.T) {
	array := [100]byte{}
	bi := NewBytesIter(array[:])
	bi.WriteUint64(100)
	bi.WriteString("aasdf")
	buf := bi.Bytes()
	copy(buf[9:], "12356")
	assert.AssertEqual(t, bi.IterUint64(), 100)
	assert.AssertEqual(t, bi.IterString(), "12356")
}

func Test2z(t *testing.T) {
	bb := NewByteBuffer(make([]byte, 6))
	bb.WriteString("今天早点做饭")
	bb.WriteString(" at ")
	t.Log(bb.len)
	s := bb.String()
	t.Log(s)
	t.Log(len(s), len("今天早点做饭 at "))
	// bb.WriteString(time.Unix(int64(1578810819), 0).Format("2006-01-02 15:04:05"))
	assert.AssertStringEqual(t, bb.String(), "今天早点做饭 at ")
}
