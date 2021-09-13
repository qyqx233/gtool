package entity

//go:generate msgp
type Person struct {
	Name    string `msg:"name"`
	Address string `msg:"-"`
	Age     int    `msg:"age"`
	Hidden  string `msg:"-"` // this field is ignored
}
