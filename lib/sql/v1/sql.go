package sql

import (
	"fmt"
	"strconv"
	"strings"
)

type KeyValuePair struct {
	Key   string
	Value interface{}
}

const (
	EqKind = iota
	NeqKind
	InKind
	NinKind
	AndKind = 1 << 32
	OrKind  = 1 << 33
)

type Node struct {
	Kind  int
	Key   string
	Value interface{}
	nodes []Node
}

func (n Node) Len() int {
	return len(n.nodes)
}

func Eqa(key string, value interface{}) Node {
	return Node{Kind: EqKind, Key: key, Value: value}
}

//go:nosplit
func Anda(ops ...Node) Node {
	var rop = Node{Kind: AndKind}
	copy(rop.nodes, ops)
	// (&rop).nodes = ops
	return rop
}

func Ora(ops ...Node) Node {
	var rop = Node{Kind: OrKind}
	// rop.nodes = ops
	copy(rop.nodes, ops)
	// for _, op := range ops {
	// 	rop.nodes = append(rop.nodes, op)
	// }
	// rop.nodes = append(rop.nodes, ops...)
	// rop.nodes = append(rop.nodes, ops...)
	return rop
}

type Op struct {
	Kind    int
	Key     string
	Value   interface{}
	as      []string
	values  []interface{}
	builder *strings.Builder
}

type AndOp struct {
}

func Eq(key string, value interface{}) Op {
	return Op{Key: key, Value: value, Kind: EqKind}
}

func Neq(key string, value interface{}) Op {
	return Op{Key: key, Value: value, Kind: NeqKind}
}

func In(key string, values interface{}) Op {
	return Op{Key: key, Value: values, Kind: InKind}
}

func Nin(key string, values interface{}) Op {
	return Op{Key: key, Value: values, Kind: NinKind}
}

func JoinInt(v []int) string {
	var as []string
	for x := range v {
		as = append(as, strconv.Itoa(x))
	}
	return strings.Join(as, ", ")
}

// func (o Op) gen() []string {
// 	if len(o.buf) != 0 {
// 		return o.buf
// 	}
// 	switch o.Kind {
// 	case EqKind:
// 		o.buf = []string{"(" + o.Key + " = ?)"}
// 	case NeqKind:
// 		o.buf = []string{"(" + o.Key + " != ?)"}
// 	case InKind:
// 		switch vv := o.Value.(type) {
// 		case []int:
// 			o.buf = []string{"(" + o.Key + " in (" + joinInt(vv) + "))"}
// 		case []string:
// 			o.buf = []string{"(" + o.Key + " in (" + strings.Join(vv, ", ") + "))"}
// 		}
// 	case NinKind:
// 		switch vv := o.Value.(type) {
// 		case []int:
// 			o.buf = []string{"(" + o.Key + " not in (" + joinInt(vv) + "))"}
// 		case []string:
// 			o.buf = []string{"(" + o.Key + " not in (" + strings.Join(vv, ", ") + "))"}
// 		}
// 	default:
// 		o.buf = []string{}
// 	}
// 	return o.buf
// }

func (o Op) string() (string, interface{}) {
	switch o.Kind {
	case EqKind:
		return "(" + o.Key + " = ?)", o.Value
	case NeqKind:
		return "(" + o.Key + " != ?)", o.Value
	case InKind:
		switch vv := o.Value.(type) {
		case []int:
			return "(" + o.Key + " in (" + JoinInt(vv) + "))", o.Value
		case []string:
			return "(" + o.Key + " in (" + strings.Join(vv, ", ") + "))", o.Value
		}
	case NinKind:
		switch vv := o.Value.(type) {
		case []int:
			return "(" + o.Key + " not in (" + JoinInt(vv) + "))", o.Value
		case []string:
			return "(" + o.Key + " not in (" + strings.Join(vv, ", ") + "))", o.Value
		}
	}
	return "1 = 1", nil
}

func (o Op) String() (string, []interface{}) {
	if o.Kind>>32 == 0 {
		sql, value := o.string()
		return sql, []interface{}{value}
	} else {
		return strings.Join(o.as, ""), o.values
	}
}

func (o Op) String1() (string, []interface{}) {
	if o.Kind>>32 == 0 {
		sql, value := o.string()
		return sql, []interface{}{value}
	} else {
		return o.builder.String(), o.values
	}
}

// func x(vs ...interface{}) interface{} {
// 	fmt.Println("x begin")
// 	var vv interface{}
// 	for _, v := range vs {
// 		fmt.Println(v)
// 		vv = v
// 	}
// 	return vv
// }

// func y(v interface{}) interface{} {
// 	fmt.Println("y ", v)
// 	return v
// }

func And(os ...Op) Op {
	var rop = Op{Kind: AndKind}
	rop.as = append(rop.as, "(")
	for _, o := range os {
		if o.Kind>>32 == 0 {
			sql, value := o.string()
			if value != nil {
				rop.values = append(rop.values, value)
			}
			rop.as = append(rop.as, sql, " and ")
		} else {
			rop.as = append(rop.as, o.as...)
			rop.as = append(rop.as, " and ")
			rop.values = append(rop.values, o.values...)
		}
	}
	rop.as = rop.as[:len(rop.as)-1]
	rop.as = append(rop.as, ")")
	return rop
}

func And1(ops ...Op) Op {
	var rop = Op{Kind: AndKind}
	op := ops[0]
	if op.builder == nil {
		op.builder = new(strings.Builder)
	}
	rop.builder = op.builder
	rop.builder.WriteString("(")
	var n = len(ops) - 1
	for i, o := range ops {
		if o.Kind>>32 == 0 {
			sql, value := o.string()
			if value != nil {
				rop.values = append(rop.values, value)
			}
			// rop.as = append(rop.as, sql, " and ")
			if i == n {
				op.builder.WriteString(sql)
			} else {
				op.builder.WriteString(sql + " and ")
			}
		} else {
			if i != n {
				op.builder.WriteString(" or ")
			}
		}
	}
	rop.builder.WriteString(")")
	return rop
}

func Or1(ops ...Op) Op {
	var rop = Op{Kind: AndKind}
	op := ops[0]
	var n = len(ops) - 1
	if op.builder == nil {
		op.builder = new(strings.Builder)
	}
	rop.builder = op.builder
	rop.builder.WriteString("(")
	for i, o := range ops {
		fmt.Printf("or %d %d\n", i, n)
		if o.Kind>>32 == 0 {
			sql, value := o.string()
			if value != nil {
				rop.values = append(rop.values, value)
			}
			fmt.Printf("or.. %d %d\n", i, n)
			if i == n {
				op.builder.WriteString(sql)
			} else {
				op.builder.WriteString(sql + " or ")
			}
		} else {
			if i != n {
				op.builder.WriteString(" or ")
			}
		}
	}
	rop.builder.WriteString(")")
	return rop
}

func Or(os ...Op) Op {
	var rop = Op{Kind: OrKind}
	rop.as = append(rop.as, "(")
	for _, o := range os {
		if o.Kind>>32 == 0 {
			sql, value := o.string()
			if value != nil {
				rop.values = append(rop.values, value)
			}
			rop.as = append(rop.as, sql, " or ")
		} else {
			rop.as = append(rop.as, o.as...)
			rop.as = append(rop.as, " or ")
			rop.values = append(rop.values, o.values...)
		}
	}
	rop.as = rop.as[:len(rop.as)-1]
	rop.as = append(rop.as, ")")
	return rop
}

func aa() {
	// fmt.Println(And(Eq("a", 1), Eq("b", 2)))
	// fmt.Println(strings.Join(Or(And(Eq("a", 1), Neq("b", "a"), In("d", []int{1, 2})),
	// 	Eq("c", 3)).Strings(), ""))
	// And(Eq("a", 1), Or())
	// Op{Key: "a", Value: 1, Next: Op{Key: "b", Value: 2, Next: nil}}
	// x(y(1), y(2))
	// y()
	// y(x(y(1), y(2)))
	// fmt.Println(Or(And(Op{Key: "a", Value: 10}, Op{Key: "b", Value: "b"}), Op{Key: "c", Value: 20}).String())
	// fmt.Println(And(Op{Key: "a", Value: 10}, Op{Key: "b", Value: "b"}).String())

	fmt.Println(Or1(And1(Op{Key: "a", Value: 10}, Op{Key: "b", Value: "b"}), Op{Key: "c", Value: 20}).String1())
	fmt.Println(And1(Op{Key: "a", Value: 10}, Op{Key: "b", Value: "b"}).String1())
	// fmt.Println(Op{Key: "c"}.String())
	// fmt.Println(or(and(Op{Key: "a"}, Op{Key: "b"}), Op{Key: "c"}).String())
}

type Node1 struct {
	Kind  int
	Key   string
	Value interface{}
	nodes []*Node1
}

func And1a(ops ...*Node1) *Node1 {
	var rop = &Node1{Kind: AndKind}
	copy(rop.nodes, ops)
	return rop
}

func Or1a(ops ...*Node1) *Node1 {
	var rop = &Node1{Kind: OrKind}
	copy(rop.nodes, ops)
	return rop
}

func Eq1a(key string, value interface{}) *Node1 {
	return &Node1{Kind: EqKind, Key: key, Value: value}
}
