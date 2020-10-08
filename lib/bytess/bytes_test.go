package bytess

import (
	"encoding/binary"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"sort"
	"testing"

	"github.com/qyqx233/gtool/lib/assert"
)

type byteBuffer1 struct {

	// B is a byte buffer to use in append-like workloads.
	// See example code for details.
	B []byte
}

// Len returns the size of the byte buffer.
func (b *byteBuffer1) Len() int {
	return len(b.B)
}

func (b *byteBuffer1) WriteString(s string) (int, error) {
	b.B = append(b.B, s...)
	return len(s), nil
}

func BenchmarkC1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bs := NewByteBuffer(make([]byte, 20))
		bs.WriteString("aaaabbbb")
		_ = bs.String()
	}
}

func BenchmarkC2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bs := NewBytesIter(make([]byte, 20))
		bs.WriteString("aaaabbbb")
		_ = bs.IterString()
	}
}

func BenchmarkC3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bs := &byteBuffer1{make([]byte, 0, 30)}
		bs.WriteString("aaaabbbb")
	}
}

func Test2x(t *testing.T) {
	array := [30]byte{}
	bi := NewBytesIter(array[:])
	var u uint64 = 21
	bi.WriteString("def")
	assert.AssertEqual(t, bi.wo, 4)
	bi.WriteUint64(u)
	assert.AssertEqual(t, bi.wo, 12)
	bi.WriteString("abc")
	bi.WriteByte(12)
	bi.WriteInt(100)
	assert.AssertTrue(t, bi.IterString() == "def")
	assert.AssertTrue(t, bi.IterUint64() == 21)
	assert.AssertTrue(t, bi.IterString() == "abc")
	assert.AssertTrue(t, bi.IterByte() == 12)
	assert.AssertTrue(t, bi.IterInt() == 100)
	bi1 := (&BytesIter{}).Loads(array[:], bi.wo)
	assert.AssertTrue(t, bi1.wo == bi.wo)
	assert.AssertTrue(t, bi1.IterString() == "def")
	bi1.WriteString("xxx")
	assert.AssertTrue(t, bi1.IterUint64() == 21)
	assert.AssertTrue(t, bi1.IterString() == "abc")
	// t.Log(bi1.Bytes())
	assert.AssertTrue(t, bi1.IterByte() == 12)
	assert.AssertTrue(t, bi1.IterInt() == 100)
	assert.AssertTrue(t, bi1.IterString() == "xxx")
}

func Test2y(t *testing.T) {
	array := [100]byte{}
	bi := NewBytesIter(array[:])
	bi.WriteUint64(100)
	bi.WriteString("aasdf")
	buf := bi.Bytes()
	copy(buf[9:], "12356")
	assert.AssertEqual(t, bi.IterUint64(), 100)
	assert.AssertEqual(t, bi.IterString(), "12356")
}

func Test2z(t *testing.T) {
	bb := NewByteBuffer(make([]byte, 6))
	bb.WriteString("今天早点做饭")
	bb.WriteString(" at ")
	t.Log(bb.len)
	s := bb.String()
	t.Log(s)
	t.Log(len(s), len("今天早点做饭 at "))
	// bb.WriteString(time.Unix(int64(1578810819), 0).Format("2006-01-02 15:04:05"))
	assert.AssertStringEqual(t, bb.String(), "今天早点做饭 at ")
}

func Test3z(t *testing.T) {
	x := 1024
	t.Log(byte(x))
	t.Log(byte(x >> 8))
}

func fx(bs []byte) []byte {
	return append(bs, "abc"...)
}

func Test4a(t *testing.T) {
	var data = []int{255, 255, 255}
	pktLen := int(uint32(data[0]) | uint32(data[1])<<8 | uint32(data[2])<<16)
	t.Log(binary.LittleEndian.Uint16([]byte{1, 2}))
	t.Log(pktLen)
	t.Log(clientFoundRows, clientLongPassword)
	var buf [9]byte
	t.Log(len(buf), cap(buf))
	r := fx(buf[:0])
	t.Log(len(r), cap(r))
	t.Log(copy(buf[3:], []byte("xyz")))
	// t.Log(inf)
	// reflect.valueOf()
}

func Test_sss(t *testing.T) {
	fd, err := os.OpenFile("a.cpp", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		t.Error(err)
	}
	defer fd.Close()
	fd.WriteString(fmt.Sprintf("auto map = unordered_map<string, int>{\n"))
	// fd.WriteString("};")
}

func Test_xxd(t *testing.T) {
	fset := token.NewFileSet()
	// 解析文件，主要解析token
	f, err := parser.ParseFile(fset, "const.go", nil, parser.ParseComments)
	if err != nil {
		t.Error(err)
	}
	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("bytess", fset, []*ast.File{f}, nil)
	if err != nil {
		t.Error(err)
	}

	type S struct {
		k, v string
	}
	var sl []S
	for _, v := range f.Scope.Objects {
		if v.Kind == ast.Con {
			name := v.Name
			val := pkg.Scope().Lookup(v.Name).(*types.Const).Val()
			sl = append(sl, S{name, val.String()})
			// d := v.Decl.(*ast.ValueSpec)
			// m[pkg.Scope().Lookup(v.Name).(*types.Const).Val().String()] = d.Comment.Text()
		}
	}
	sort.Slice(sl, func(i, j int) bool {
		return sl[i].k < sl[j].k
	})
	fd, err := os.OpenFile("a.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		t.Error(err)
	}
	defer fd.Close()
	for _, v := range sl {
		fd.WriteString("#define " + v.k + " " + v.v + "\n")
	}
}
