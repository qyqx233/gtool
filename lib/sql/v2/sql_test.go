package sql

import (
	"fastweb/lib/bytess"
	"fastweb/lib/sql/v1"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

func foo(vs *[]interface{}) {
	*vs = append(*vs, 1)
}

func setM(mp map[string]string) {
	mp["a"] = "aa"
}

func setQ(q []int) {
	q[0] = 100
}

type bb struct {
	buf   []byte
	store [40]byte
}

var _buf = [100]byte{}

func newBuf(n int) *bb {
	b := new(bb)
	return b
	// return &bb{buf: _buf[:100]}
}

func (b *bb) WriteString(s string) {
	b.buf = append(b.buf, s...)
}

func (b *bb) String() string {
	return *(*string)(unsafe.Pointer(&b.buf))
}

func foo1(ss []int) {
	ss = append(ss, 1)
}

func Test5(t *testing.T) {
	var mp = make(map[string]string)
	var q = make([]int, 10)
	setQ(q)
	setM(mp)
	t.Log(mp, q)
	var b = newBuf(10)
	t.Log(reflect.TypeOf(b))
	b.WriteString("11")
	t.Log(string(b.buf))
	// t.Log(333)
	// b := newBuf(100)
	// b.WriteString("abcde")
	// t.Log(string(b.buf))
	// t.Log(b.String())
	// b = newBuf(100)
	// b.WriteString("123") z
	// t.Log(b.String())
}

func Test4(t *testing.T) {
	r := Or(And(Eq("a", 1), Eq("x", 300)), In("aa", []int{1, 2, 3}))
	sql, _ := r.Sql()
	t.Log(sql)
	r = Eq("a", 100)
	t.Log(r.Sql())
}

func BenchmarkX1(b *testing.B) {
	// w, _ := os.OpenFile("E:\\Develop\\fastweb\\lib\\sql\\profile", os.O_WRONLY|os.O_CREATE, 0600)
	// defer w.Close()
	// pprof.StartCPUProfile(w)
	// defer pprof.StopCPUProfile()
	for i := 0; i < b.N; i++ {
		r := Or(And(Eq("a", 1), Eq("x", 300)), Eq("aa", 100))
		_, _ = r.Sql()
	}
}

func BenchmarkB1(b *testing.B) {
	// w, _ := os.OpenFile("E:\\Develop\\fastweb\\lib\\sql\\profile", os.O_WRONLY|os.O_CREATE, 0600)
	// defer w.Close()
	// pprof.StartCPUProfile(w)
	// defer pprof.StopCPUProfile()
	for i := 0; i < b.N; i++ {
		buf := &bb{}
		buf.WriteString("12345678901234567890")
	}
}

func fxx() {
	r := Or(And(Eq("a", 1), Eq("x", 300)), Eq("aa", 100))
	_, _ = r.Sql()
}

func BenchmarkB2(b *testing.B) {
	// w, _ := os.OpenFile("E:\\Develop\\fastweb\\lib\\sql\\profile", os.O_WRONLY|os.O_CREATE, 0600)
	// defer w.Close()
	// pprof.StartCPUProfile(w)
	// defer pprof.StopCPUProfile()
	for i := 0; i < b.N; i++ {
		// Or(And(Eq("a", 1), Eq("x", 300)), Eq("aa", 100))
		And(Or(Eq("xx", 100), Eq("xx", 100)), Eq("xx", 100))
	}
}

func no(v interface{}) {
}

func BenchmarkB3(b *testing.B) {
	// sql.Ora(sql.Node{Kind: 1, Key: "a", Value: "b"})
	for i := 0; i < b.N; i++ {
		// Or(And(Eq("a", 1), Eq("x", 300)), Eq("aa", 100))
		// a := sql.Eqa("xx", 100)
		// sql.Ora(&a, sql.Node("xx", 100), sql.Node("xx", 100))
		// no(sql.Eqa("a", 100))
		// sql.Ora(sql.Node{Kind: 1, Key: "a", Value: "b"})
		// sql.Ora(sql.Eqa("a", "b"))
		sql.Ora(sql.Anda(sql.Ora(sql.Eqa("xx", 100)), sql.Eqa("aa", 33)), sql.Eqa("cc", 10))
	}
}

func Test6(t *testing.T) {
	v := sql.Ora(sql.Eqa("11", 1))
	t.Log(v.Len())
}

type outBuffer struct {
}

func (outBuffer) WriteString(s string) {
	fmt.Println(s)
}

type wrapNode struct {
	node  *Node
	state int
}

func travel1(node *Node, buffer *bytess.ByteBuffer) {
walk:
	if node.Kind>>32 == 0 {
		stringOp1(node, buffer)
	} else {
		n := len(node.nodes) - 1
		buffer.WriteString("(")
		for i, child := range node.nodes {
			// travel(&child, buffer, values)
			node = &child
			goto walk
			if i != n {
				// 	switch node.Kind {
				// 	case sql.AndKind:
				// 	case sql.OrKind:
				// 	}
			}
		}
	}
}

func kont(n int) {
	return
}

func fact(n int) int {
	if n == 1 {
		return 1
	}
	return 0
}

type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

type byteSlice struct {
	noCopy
	buf []byte
	len int
}

func newSlice(buf []byte) *byteSlice {
	s := byteSlice{buf: buf}
	return &s
}

func (bs *byteSlice) add(s string) {
	if bs.len+len(s) > cap(bs.buf) {
		var nb = make([]byte, 0, bs.len<<1+len(s))
		copy(nb, bs.buf)
		copy(nb[bs.len:], s)
		bs.buf = nb
	} else {
		copy(bs.buf[bs.len:], s)
	}
	bs.len += len(s)
}

func (bs *byteSlice) foo(s string) {
	fmt.Println(len(bs.buf), cap(bs.buf))
}

func (bs *byteSlice) add1(s string) {
	if bs.len+len(s) < cap(bs.buf) {
		var nb = make([]byte, 0, bs.len<<1+len(s))
		copy(nb, bs.buf)
		copy(nb[len(nb):], s)
		bs.buf = nb
	} else {
		copy(bs.buf[len(bs.buf):], s)
	}
	// bs.len = len(bs.buf)
	// bs.cap = cap(bs.buf)
}

func Test7(t *testing.T) {
	n := Or(Eq("xx", 100), Eq("xx", 100))
	// n := And(Or(Eq("xx", 100), Eq("xx", 100)), Eq("xx", 100))
	t.Log(len(n.nodes), n.Kind == sql.OrKind)
	t.Log(n.Sql1())
}

type ByteSlice []byte

func (b *ByteSlice) add(s string) ByteSlice {
	var bb = *b
	if len(bb)+len(s) < len(bb) {
		x := make([]byte, len(bb)*2+len(s))
		copy(x, bb)
		copy(x[len(bb):], s)
		*b = x
	}
	copy(bb[len(bb):], s)
	return bb
}

func (b *ByteSlice) WriteString(s string) {
	*b = append(*b, s...)
}

func newSlice1(b []byte) ByteSlice {
	return b
}

var long = strings.Repeat("a", 110)

func BenchmarkC1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := [100]byte{}
		bs := newSlice(buf[:])
		bs.add("long")
	}
}

func fx(a *[]int) {
	// *a = append(*a, 1)
}

func fy(bs *ByteSlice) {
	bs.WriteString("abc")
}

func (n *Node) Or(ops ...Node) *Node {
	n.Kind = sql.OrKind
	n.nodes = append(n.nodes, ops...)
	return n
}

func Or1(ops ...Node) Node {
	var rop = Node{Kind: sql.OrKind}
	rop.nodes = append(rop.nodes, ops...)
	return rop
}

type Nodes []Node

func New() Node {
	var n = Node{}
	var an = [8]Node{}
	n.nodes = an[:0:8]
	return n
}

func (n Node) Or1(ops ...Node) Node {
	n.nodes = append(n.nodes, ops...)
	return n
}

func BenchmarkC2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// new(Node).Or(Eq("a", 1), Eq("b", "1234"))
		_ = New().Or1(Eq("a", 1), Eq("b", "1234"))

		// l := []int{1, 3, 4}
		// var s = l[:0:3]
		// s = append(s, 1, 2, 3, 4)
	}
}

func Test8(t *testing.T) {
	// n := New().Or1(Eq("x", 100))
	// t.Log(len(n), cap(n), n[0].Key)
	l := []int{1, 3, 4}
	var s = l[:0:3]
	t.Log(len(s), cap(s))
}
