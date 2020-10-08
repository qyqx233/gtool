package test

import (
	"testing"
)

type sx struct {
	a, b  int
	s, s1 string
}

func newSx() *sx {
	return &sx{}
}

func newSx1(x *sx) *sx {
	return x
}

func BenchmarkA1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newSx()
	}
}

func BenchmarkA2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := sx{}
		newSx1(&x)
	}
}
