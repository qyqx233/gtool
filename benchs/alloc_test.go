package benches

import (
	"fmt"
	"strings"
	"testing"
)

type goodSlice struct {
	buf []int
}

func modifySlice(sl []Op) {
	copy(sl, []Op{Op{}})
}

func modifySlice1(sl *[]Op) {
	*sl = append(*sl, Op{})
}

func modifySlice2(sl *[]Op) {
	var ss = *sl
	fmt.Println(len(ss), cap(ss))
	if len(ss) < cap(ss) {
		ss[len(ss)] = Op{}
	}
}

func copyS(sl []Op) {
	copy(sl, []Op{Op{"a", 100}})
}

func TestA(t *testing.T) {
	// a := [1]int{}
	// copy(a[:], []int{1})
	// t.Log(a)

	os := [3]Op{}
	copyS(os[:])
	t.Log(os)
}

func TestA2(t *testing.T) {
	var s = [3]int{}
	var ss = s[:]
	t.Log(cap(ss), len(ss))
	copy(ss, []int{1})
	t.Log(cap(ss), len(ss))
}

func TestA3(t *testing.T) {
	var sl = make([]Op, 0, 3)
	modifySlice2(&sl)
}

func BenchmarkA1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sl := [3]Op{}
		modifySlice((sl[:]))
	}
}

func BenchmarkA2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make([]Op, 0, 4)
		var sl = make([]Op, 0, 4)
		sl = append(sl, Op{})
		// modifySlice1(&sl)
	}
}
func BenchmarkA3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		builder.Grow(40)
		builder.WriteString("1234")
		// modifySlice1(&sl)
	}
}
