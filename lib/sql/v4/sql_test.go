package sql

import (
	"fastweb/lib/bytess"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unsafe"

	"github.com/valyala/bytebufferpool"
)

type ByteSlice []byte

func (b *ByteSlice) WriteString(s string) {
	copy(*b, s)
}

var long = strings.Repeat("a", 101)

func BenchmarkC1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var y = Mul{And: []Mul{Eq("abc"), Eq("def"), Mul{Or: []Mul{Eq("x"), Eq("y")}}}}
		vs := bytess.NewValueBuffer(make([]interface{}, 10))
		bs := bytess.NewByteBuffer(make([]byte, 100))
		y.string(bs, vs)
	}
}

func BenchmarkCx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var y = Mul{And: []Mul{Eq("abc"), Eq("def"), Mul{Or: []Mul{Eq("x"), Eq("y")}}}}
		display1(&y)
	}
}

type AA struct {
	a int
}

func (a *AA) set() {
	a.a = 100
}

func display(m *Mul) {
	var p = &(m.And[2].Or[0].Key)
	fmt.Println(reflect.TypeOf(p).Kind())
	// reflect.ValueOf(p)
	fmt.Println("=>", reflect.TypeOf(&(m.And[2].Or[0].Key)))
	switch reflect.TypeOf(&(m.And[2].Or[0].Key)).String() {
	case "*string":

	}
}

func display1(m *Mul) {
	// _ = reflect.TypeOf(&(m.And[2].Or[0].Key))
	var p = &(m.And[2].Or[0].Key)
	reflect.ValueOf(p)
}

func Test9(t *testing.T) {
	var y = Mul{And: []Mul{Eq("abc"), Eq("def"), Mul{Or: []Mul{Eq("x"), Eq("y")}}}}
	t.Log(y.And[2].Or[0].Key)
	t.Logf("%p %p\n", &y, &(y.And[2].Or[0].Key))
	var p *string = &(y.And[2].Or[0].Key)
	t.Log(*p)
	display(&y)
	display(&y)
	vs := bytess.NewValueBuffer(make([]interface{}, 10))
	bs := bytess.NewByteBuffer(make([]byte, 100))
	y.string(bs, vs)
	t.Log(bs.String())
}

type iface [2]uintptr

func Test10(t *testing.T) {
	var s = "abcd"
	var v, v1 interface{}
	var pi *uint8
	var pj *uint
	v = s
	// 使用一个interface{}类型的变量指向一个对象（这里是个字符串）的时候， 【type,unsafe.Pointer】,
	// pointer 并不指向实际的地址，解引用后得到interface对象引用对象的实际地址
	//
	v1 = s
	var p, p1 *iface
	t.Log(v, v1, unsafe.Sizeof(v))
	p = (*iface)(unsafe.Pointer(&s))
	t.Log(p)
	pi = (*uint8)(unsafe.Pointer(p[0]))
	t.Log("pi=", *pi)
	p = (*iface)(unsafe.Pointer(&v))
	pi = (*uint8)(unsafe.Pointer(p[1]))
	pj = (*uint)(unsafe.Pointer(p[1]))
	t.Log(p, &v, *pj)
	p1 = (*iface)(unsafe.Pointer(&v1))
	pj = (*uint)(unsafe.Pointer(p1[1]))
	t.Log(p1, &v1, *pj)

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(p[i], p1[i])
	// }
	var ptr = uintptr((unsafe.Pointer(p[0])))
	ptr++
	t.Log(*(*uint8)(unsafe.Pointer(ptr)))
}

func Test11(t *testing.T) {

	var ss = []byte("1234")
	var s = string(ss)
	var p = (*iface)(unsafe.Pointer(&s))
	t.Log(p)
	// pi = (*uint8)(unsafe.Pointer(p[0]))
	var base = p[0]
	for i := 0; i < len(s); i++ {
		pi := (*uint8)(unsafe.Pointer(base + uintptr(i)))
		*pi = 100
		t.Log(*pi)
	}
}

func foo(b *bytebufferpool.ByteBuffer) {
	b.WriteString("aaaa")
}

func BenchmarkD1(b *testing.B) {
	buffer := bytebufferpool.ByteBuffer{}
	foo(&buffer)
}

type bs struct {
	buf []byte
}

func (b *bs) WriteString(s string) {
	b.buf = append(b.buf, s...)
}

func goo(b *bs) {
	b.WriteString("123")
}

func xoo(b *bytess.ByteBuffer) {
	b.WriteString(long)
}

func BenchmarkD2(b *testing.B) {
	buffer := bs{}
	// buffer.WriteString("asadfqwe")
	goo(&buffer)
}

func BenchmarkD3(b *testing.B) {
	buffer := bytess.NewByteBuffer(make([]byte, 8))
	// buffer.WriteString("asadfqwe")
	xoo(buffer)
}
