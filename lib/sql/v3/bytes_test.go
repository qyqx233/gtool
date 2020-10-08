package sql

import (
	"fastweb/lib/bytess"
	"fmt"
	"testing"
)

func and(vs ...interface{}) interface{} {
	l := len(vs)
	for i := 0; i < l; i++ {
		if i != 0 {
			fmt.Print(" and ")
		}
	}
	return vs[0]
}

func or(vs ...interface{}) interface{} {
	// fmt.Println("or")
	// return vs[0]
	l := len(vs)
	for i := 0; i < l; i++ {
		if i != 0 {
			fmt.Print(" or ")
		}
	}
	return vs[0]
}

func quote(v ...interface{}) interface{} {
	return v
}

type deque struct {
	len, cap int
	buf      []byte
}

func (dq *deque) insert(s string) {
	copy(dq.buf[dq.len:], s)
}

type Nx struct {
	Key string
	dq  *bytess.Deque
}

func fx(s1, s2 string) string {
	return s1 + s2
}

func fy(s1, s2 string) string {
	return s1 + s2
}

func (x Nx) execute() {
	x.dq.Append(x.Key)
}

func BenchmarkX1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fx(fx(fy("1", "2"), "3"), "4")
	}
}

func Test3(t *testing.T) {

}

func yy(dq deque) {
}
