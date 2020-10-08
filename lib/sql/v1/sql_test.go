package sql

import (
	"bytes"
	"os"
	"runtime/pprof"
	"strings"
	"testing"
)

func Test2(t *testing.T) {
	// var ss []interface{}
	// ss = append(ss, 1)
	// t.Log(ss)
	// var s = []string{"1"}
	// var in interface{}
	// in = s
	// switch v := in.(type) {
	// case []int:
	// 	t.Log(v)
	// case []string:
	// 	t.Log(v)
	// }
}

func Test1(t *testing.T) {
	aa()
	// t.Log(EqKind, OrKind)
	// var s = []string{"a", "b", "c"}
	// s = s[:len(s)-1]
	// t.Log(s)
}

func Test4(t *testing.T) {
	// y(x(y(1), y(2)))
}

func Test3(t *testing.T) {
	// var i uint
	// i = 1 << 32
	// t.Log(fmt.Sprintf("%x", i))
	// t.Log(fmt.Sprintf("%x", i>>32))
	// t.Log(OrKind>>32, AndKind>>32, 1<<32|1, 1<<32)
}

func countOps(ops ...Op) int {
	var n int
	for _ = range ops {
		n++
	}
	return n
}

func addStr(s1, s2, s3 string) string {
	var buffer bytes.Buffer
	buffer.WriteString(s1)
	buffer.WriteString(s2)
	return buffer.String()
}

func addStr2(s1, s2, s3 string) string {
	var buffer strings.Builder
	buffer.WriteString(s1)
	buffer.WriteString(s2)
	return buffer.String()
}

func addStr1(s1, s2, s3 string) string {
	return s1 + s2
}

var s1 = "asdfasaaa"
var s2 = "sdfasaaa"
var s3 = "sdfasaa"

func BenchmarkA5(b *testing.B) {

	for i := 0; i < b.N; i++ {
		addStr1(s1, s2, s3)
	}
}

func BenchmarkA6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addStr2(s1, s2, s3)
	}
}

func BenchmarkA4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addStr(s1, s2, s3)
	}
}

func BenchmarkA3(b *testing.B) {
	var ops = []Op{Op{}, Op{}, Op{}}
	for i := 0; i < b.N; i++ {
		countOps(ops...)
	}
}

func BenchmarkA2(b *testing.B) {
	w, _ := os.OpenFile("E:\\Develop\\fastweb\\lib\\sql\\profile", os.O_WRONLY|os.O_CREATE, 0600)
	defer w.Close()
	pprof.StartCPUProfile(w)
	defer pprof.StopCPUProfile()
	for i := 0; i < b.N; i++ {
		And(Op{Key: "a", Value: 10}, Op{Key: "b", Value: "b"})
	}
	defer pprof.StopCPUProfile()
}

func BenchmarkX1(b *testing.B) {
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
		Or1a(And1a(Or1a(Eq1a("xx", 100)), Eq1a("aa", 33)), Eq1a("cc", 10))
	}
}
