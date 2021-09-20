package cache

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/qyqx233/gtool/cache/entity"
)

func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

var buff = make([]byte, 0, 16)

func BenchmarkCacheGet1(b *testing.B) {
	const items = 1 << 16
	c := New(12 * items)
	defer c.Reset()
	k := []byte("\x00\x00\x00\x00")
	v := []byte("xyza")
	for i := 0; i < items; i++ {
		k[0]++
		if k[0] == 0 {
			k[1]++
		}
		c.Set(k, v)
	}

	for i := 0; i < b.N; i++ {
		var buf []byte
		k := []byte("\x98\x53\x00\x00")
		buf = c.Get(buf, k)
		if string(buf) != string(v) {
			panic(fmt.Errorf("BUG: invalid value obtained; got %q; want %q", buf, v))
		}
	}

}

func BenchmarkCacheGet(b *testing.B) {
	const items = 1 << 16
	c := New(12 * items)
	defer c.Reset()
	k := []byte("\x00\x00\x00\x00")
	v := []byte("xyza")
	for i := 0; i < items; i++ {
		k[0]++
		if k[0] == 0 {
			k[1]++
		}
		c.Set(k, v)
	}

	b.ReportAllocs()
	b.SetBytes(items)
	b.RunParallel(func(pb *testing.PB) {
		var buf []byte
		k := []byte("\x00\x00\x00\x00")
		for pb.Next() {
			for i := 0; i < items; i++ {
				k[0]++
				if k[0] == 0 {
					k[1]++
				}
				buf = c.Get(buf[:0], k)
				if string(buf) != string(v) {
					panic(fmt.Errorf("BUG: invalid value obtained; got %q; want %q", buf, v))
				}
			}
		}
	})
}

func Test1(t *testing.T) {
	const items = 1 << 16
	c := New(12 * items)
	defer c.Reset()
	k := []byte("\x00\x00\x00\x00")
	v := []byte("xyza")
	c.Set(k, v)
}

func add(b []byte) []byte {
	b = append(b, 10)
	return b
}

func Benchmark2(b *testing.B) {
	var buf = make([]byte, 0, 12)
	for i := 0; i < b.N; i++ {
		add(buf)
	}
}

func Get(p Marshaler, s string) error {
	return nil
}

func Test2(t *testing.T) {
	var out []byte
	var err error
	const items = 1 << 16
	c := New(12 * items)
	cacher := Cacher{c}
	get := cacher.GetString(Get)
	var p = new(entity.Person)
	p.Name = "hello"
	out, _ = p.MarshalMsg(out)
	c.Set([]byte("ss"), out)
	var dst []byte
	c.Get(dst, []byte("ss"))
	var p1 = new(entity.Person)
	err = get(p1, "ss")
	if err != nil {
		t.Error(err)
	}
	t.Log(p1)
}

func Benchmark_all(b *testing.B) {
	var out []byte
	const items = 1 << 16
	var p = new(entity.Person)
	p.Name = "hello"
	p.Address = "asdfqwer"
	out, _ = p.MarshalMsg(out)
	c := New(12 * items)
	c.Set([]byte("ss"), out)
	cacher := Cacher{c}
	get := cacher.GetString(Get)
	var p1 = new(entity.Person)
	for i := 0; i < b.N; i++ {
		_ = get(p1, "ss")
	}
	b.Log(p)
}

func cacher() *Cacher {
	const items = 1 << 16
	c := New(12 * items)
	return &Cacher{c}
}

func echoInt(p Marshaler, s uint64) error {
	fmt.Println("echoInt")
	p.(*entity.Person).Age = 10
	return nil
}

func echo(p Marshaler, s string) error {
	p.(*entity.Person).Age = 10
	return nil

}
func Test_int_deploy(b *testing.T) {
	var p entity.Person
	c := cacher()
	cacheEcho := c.GetIntPrefix(echoInt, "p:")
	cacheEcho(&p, 1)
	cacheEcho(&p, 1)
	cacheEcho(&p, 2)
}
func Test_deploy(b *testing.T) {
	var p entity.Person
	c := cacher()
	cacheEcho := c.GetStringPrefix(echo, "p:")
	cacheEcho(&p, "a")
	cacheEcho(&p, "a")
}
func Benchmark_int_deploy(b *testing.B) {
	c := cacher()
	var p entity.Person
	cacheEcho := c.GetIntPrefix(echoInt, "p:")
	for i := 0; i < b.N; i++ {
		cacheEcho(&p, 1)
	}
}

func Benchmark_deploy(b *testing.B) {
	c := cacher()
	var p entity.Person
	cacheEcho := c.GetStringPrefix(echo, "p:")
	for i := 0; i < b.N; i++ {
		cacheEcho(&p, "a")
	}
}
