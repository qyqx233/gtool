package datastruct

import (
	"testing"
)

type A struct {
	s string
}

func BenchmarkC2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack := NewStack()
		stack = stack.Push(A{})
		v, stack := stack.Pop()
		_ = v.(A)
	}
}

func Test1(t *testing.T) {
	// var arr = Stack([]int{})
	// arr.Push()
	stack := NewStack()
	stack = stack.Push(A{})
	t.Log(len(stack))
	_, stack = stack.Pop()
}
