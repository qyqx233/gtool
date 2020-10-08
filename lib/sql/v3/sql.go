package sql

import (
	"fastweb/lib/bytess"
	"fastweb/lib/sql/v1"
)

type Node struct {
	Kind     int
	Key      string
	Value    interface{}
	nodes    []Node
	len, cap int
}

func Logic(kind int, nodes ...Node) Node {
	switch kind {
	}
	n := Node{Kind: kind, nodes: make([]Node, 8)}
	copy(n.nodes, nodes)
	return n
}

func AndNodes(nodes ...Node) Node {
	return Logic(sql.AndKind, nodes...)
}

type Op interface {
	string(*bytess.Deque)
}

type Eqs struct {
	Key   string
	Value interface{}
}

func (v Eqs) string(dq *bytess.Deque) {
	dq.Append(v.Key)
}

type Ands struct {
	// Key   string
	// Value interface{}
	Ops []Op
}

type Mul struct {
	Key   string
	Value interface{}
	And   []Mul
	Or    []Mul
}

type Mul1 struct {
	Key   string
	Value interface{}
	Ops   [10]*Mul1
}

func Add(m ...*Mul1) Mul1 {
	x := Mul1{}
	copy(x.Ops[:], m)
	return x
}

func newMul() {

}

func dump(v Op, dq *bytess.Deque) {
	// dq := bytess.NewDeque(make([]byte, 60))
	v.string(dq)
	// switch vv := v.(type) {
	// case Ands:
	// 	v.
	// case Eqs:
	// 	dq.Append(vv.string())
	// }

}

func (v Ands) string(dq *bytess.Deque) {
	for _, n := range v.Ops {
		n.string(dq)
	}
}

func Eq(key string, value interface{}) Node {
	return Node{Kind: sql.EqKind, Key: key, Value: value}
}

type Builder struct {
	dq *bytess.Deque
}

func NewBuilder(bs []byte) *Builder {
	return &Builder{dq: bytess.NewDeque(bs)}
}

func (b *Builder) And(ops ...Node) Node {
	l := len(ops)
	b.dq.Append("(")
	for i := 0; i < l; i++ {
		if i != 0 {
			b.dq.Append(" and ")
		}
		op := ops[i]
		if op.Kind>>32 == 0 {
			b.dq.Append(op.Key)
			b.dq.Append(" = ")
		}
	}
	b.dq.Append(")")
	return Node{Kind: sql.AndKind}
}

func (b *Builder) Or(ops ...Node) Node {
	l := len(ops)
	b.dq.Append("(")
	for i := 0; i < l; i++ {
		if i != 0 {
			b.dq.Append(" or ")
		}
		op := ops[i]
		if op.Kind>>32 == 0 {
			b.dq.Append(op.Key)
			b.dq.Append(" = ")
		}
	}
	b.dq.Append(")")
	return Node{Kind: sql.OrKind}
}
