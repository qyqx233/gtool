package main

import (
	"crypto/sha256"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"log"
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func b2Hex(b []byte) string {
	var ss = "0123456789abcdef"
	bb := make([]byte, len(b)*2)
	for i, n := range b {
		bb[2*i] = ss[n>>4&0x0f]
		bb[2*i+1] = ss[n&0x0f]
	}
	return string(bb)
}

func scrambleSHA256Password(scramble []byte, password string) []byte {
	if len(password) == 0 {
		return nil
	}
	crypt := sha256.New()
	crypt.Write([]byte(password))
	message1 := crypt.Sum(nil)
	crypt.Reset()
	crypt.Write(message1)
	message1Hash := crypt.Sum(nil)
	crypt.Reset()
	crypt.Write(message1Hash)
	crypt.Write(scramble)
	message2 := crypt.Sum(nil)
	for i := range message1 {
		message1[i] ^= message2[i]
	}
	return message1
}

func Test3(t *testing.T) {
	var s = []byte{'a', 'b', 'c'}
	crypt := sha256.New()
	crypt.Write(s)
	message1 := crypt.Sum(nil)
	t.Log(b2Hex(message1), len(message1))
	t.Log(b2Hex(scrambleSHA256Password([]byte("abcdef"), "123456")))
}

func Test1(t *testing.T) {
	pkgInfo, err := build.ImportDir(".", 0)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range pkgInfo.GoFiles {
		if file != "a.go" {
			continue
		}
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, file, nil, 0)
		if err != nil {
			log.Fatal(err)
		}
		ast.Inspect(f, func(n ast.Node) bool {
			switch v := n.(type) {
			case *ast.StructType:
				t.Log("StructType")
				for _, f := range v.Fields.List {
					t.Log(f.Names, reflect.TypeOf(f.Type))
					switch ff := f.Type.(type) {
					case *ast.ArrayType:
						t.Log(ff.Pos(), ff.End(), ff.Elt.(*ast.Ident).Name, reflect.TypeOf(ff.Elt))
					}
				}
			case *ast.GenDecl:
				// for _, spec := range v.Specs {
				// t.Log(spec.Pos(), spec.End())
				// switch x := spec.(type) {
				// case *ast.ImportSpec:
				// case *ast.TypeSpec:
				// 	t.Log(x.Name, x.Type.Pos(), x.Type.End())
				// case *ast.ValueSpec:
				// }
				// }
			case *ast.FuncDecl:
				// arg0 := v.Type.Params.List[0]
				// t.Log(arg0.Names[0], arg0.Type, arg0.Tag, arg0.Comment)
				// for i := 0; i < v.Type.Params.NumFields(); i++ {
				// 	arg := v.Type.Params.List[i]
				// 	t.Log(arg.Names[0], arg.Type)
				// }

			case nil:
			}
			return true
		})
	}
}
