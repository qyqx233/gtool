package benches

import (
	"testing"
)

type Op struct {
	s string
	a int
}

func iter1(ops ...Op) int {
	var n int
	for _ = range ops {
		n++
	}
	return n
}

func iter1x(ops ...Op) int {
	var n int
	for _ = range ops[:len(ops)-1] {
		n++
	}
	n++
	return n
}

func iter1y(ops ...Op) int {
	var n int
	var x = len(ops) - 1
	for i, _ := range ops {
		if i == x {
			n++
		} else {
			n++
		}
	}
	n++
	return n
}

func BenchmarkIter11(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iter1(Op{}, Op{}, Op{})
	}
}

func BenchmarkIter11x(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iter1x(Op{}, Op{}, Op{})
	}
}

func BenchmarkIter11y(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iter1x(Op{}, Op{}, Op{})
	}
}

func iter2(ops []Op) int {
	var n int
	for i := 0; i < len(ops); i++ {
		n++
	}
	return n
}

func BenchmarkIter1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iter1(Op{}, Op{}, Op{})
	}
}

func BenchmarkIter1_a(b *testing.B) {
	ops := []Op{Op{}, Op{}, Op{}}
	for i := 0; i < b.N; i++ {
		iter1(ops...)
	}
}

func BenchmarkIter2(b *testing.B) {
	ops := []Op{Op{}, Op{}, Op{}}
	for i := 0; i < b.N; i++ {
		iter2(ops)
	}
}
