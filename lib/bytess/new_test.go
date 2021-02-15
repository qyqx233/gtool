package bytess

import "testing"

func TestX2(t *testing.T) {
	bs := NewByteBuffer(make([]byte, 5))
	bs.WriteString("abc")
	if bs.String() != "abc" || bs.Len() != 3 || bs.Cap() != 5 {
		t.Fail()
		return
	}
	bs.WriteString("111")
	if bs.String() != "abc111" || bs.Len() != 6 || bs.Cap() != 9 {
		t.Fail()
		return
	}
	bs = NewByteBuffer(make([]byte, 5))
	bs.WriteBytes([]byte("abc"))
	if bs.String() != "abc" || bs.Len() != 3 || bs.Cap() != 5 {
		t.Error("bb")
		t.Fail()
		return
	}
	bs.WriteBytes([]byte("111"))
	if bs.String() != "abc111" || bs.Len() != 6 || bs.Cap() != 9 {
		t.Error("aa")
		t.Fail()
		return
	}
}

func BenchmarkX1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bs := NewByteBuffer(make([]byte, 20))
		bs.WriteString("aaaabbbb")
		_ = bs.String()
	}
}
