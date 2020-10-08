package main

// import (
// 	"bytes"
// 	"fastweb/lib/bytess"
// 	"os"
// 	"testing"
// 	"text/template"
// )

// var bs []byte

// func init() {
// 	bs = bytes.Repeat([]byte("a"), 4096)
// }

// func BenchmarkTest(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		var buf bytess.Buffer
// 		(&buf).Init1(make([]byte, 0, 400))
// 		// buff := (&bytess.Buffer{}).Init(4096)
// 		// buff.Init(4096)
// 		buf.Write(bs)
// 	}
// }

// type sx struct {
// 	a int
// }

// func newSx() *sx {
// 	return &sx{}
// }

// func Test4(t *testing.T) {
// 	x := newSx()
// 	println(x)
// }

// func BenchmarkTest1(b *testing.B) {
// 	bb := make([]byte, 4000)
// 	for i := 0; i < b.N; i++ {
// 		buff := bytes.NewBuffer(bb)
// 		buff.Write(bs)
// 	}
// }

// func Test2(t *testing.T) {
// 	var s = `<html>
//     <head>
//         <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
//         <title>Go Web</title>
//     </head>
//     <body>
//         {{ .Age }}
//     </body>
// </html>
// 	`
// 	tmpl, err := template.New("").Parse(s)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	type a struct {
// 		Age int
// 	}
// 	tmpl.Execute(os.Stdout, map[string]string{"A": "100"})
// 	t.Log(tmpl)
// }
