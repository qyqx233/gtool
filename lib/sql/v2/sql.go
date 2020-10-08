package sql

import (
	"fastweb/lib/bytess"
	"fastweb/lib/datastruct"
	"fastweb/lib/helper"
	"fastweb/lib/sql/v1"
	"fmt"
	"strings"
)

type Node struct {
	Kind  int
	Key   string
	Value interface{}
	nodes []Node
}

//go:nosplit
func And(ops ...Node) Node {
	var rop = Node{Kind: sql.AndKind}
	copy(rop.nodes, ops)
	return rop
}

func Or(ops ...Node) Node {
	var rop = Node{Kind: sql.OrKind}
	copy(rop.nodes, ops)
	return rop
}

//go:nosplit
func Eq(key string, value interface{}) Node {
	return Node{Kind: sql.EqKind, Key: key, Value: value}
}

func Neq(key string, value interface{}) Node {
	return Node{Kind: sql.NeqKind, Key: key, Value: value}
}

func In(key string, value interface{}) Node {
	return Node{Kind: sql.InKind, Key: key, Value: value}
}

func Nin(key string, value interface{}) Node {
	return Node{Kind: sql.NinKind, Key: key, Value: value}
}

func stringOp1(o *Node, buffer *bytess.ByteBuffer) interface{} {
	switch o.Kind {
	case sql.EqKind:
		buffer.WriteString(o.Key)
		buffer.WriteString(" = ?")
	case sql.NeqKind:
		buffer.WriteString(o.Key)
		buffer.WriteString(" != ?")
	case sql.InKind:
		switch vv := o.Value.(type) {
		case []int:
			buffer.WriteString(o.Key)
			buffer.WriteString(" in (")
			buffer.WriteString(sql.JoinInt(vv))
			buffer.WriteString(")")
		case []string:
			buffer.WriteString(o.Key)
			buffer.WriteString(" in (")
			buffer.WriteString(strings.Join(vv, ", "))
			buffer.WriteString(")")
		}
	case sql.NinKind:
		switch vv := o.Value.(type) {
		case []int:
			buffer.WriteString(o.Key)
			buffer.WriteString(" not in (")
			buffer.WriteString(sql.JoinInt(vv))
			buffer.WriteString(")")
		case []string:
			buffer.WriteString(o.Key)
			buffer.WriteString(" not in (")
			buffer.WriteString(strings.Join(vv, ", "))
			buffer.WriteString(")")
		}
	}
	if o.Value != nil {
		return o.Value
	}
	return nil
}

func stringOp(o *Node, buffer *strings.Builder, values *[]interface{}) {
	switch o.Kind {
	case sql.EqKind:
		buffer.WriteString(o.Key)
		buffer.WriteString(" = ?")
		if o.Value != nil {
			*values = append(*values, o.Value)
		}
	case sql.NeqKind:
		buffer.WriteString(o.Key)
		buffer.WriteString(" != ?")
		if o.Value != nil {
			*values = append(*values, o.Value)
		}
	case sql.InKind:
		switch vv := o.Value.(type) {
		case []int:
			buffer.WriteString(o.Key)
			buffer.WriteString(" in (")
			buffer.WriteString(sql.JoinInt(vv))
			buffer.WriteString(")")
		case []string:
			buffer.WriteString(o.Key)
			buffer.WriteString(" in (")
			buffer.WriteString(strings.Join(vv, ", "))
			buffer.WriteString(")")
		}
	case sql.NinKind:
		switch vv := o.Value.(type) {
		case []int:
			buffer.WriteString(o.Key)
			buffer.WriteString(" not in (")
			buffer.WriteString(sql.JoinInt(vv))
			buffer.WriteString(")")
		case []string:
			buffer.WriteString(o.Key)
			buffer.WriteString(" not in (")
			buffer.WriteString(strings.Join(vv, ", "))
			buffer.WriteString(")")
		}
	}
}

func travel(node *Node, buffer *strings.Builder, values *[]interface{}) {
	if node.Kind>>32 == 0 {
		stringOp(node, buffer, values)
	} else {
		n := len(node.nodes) - 1
		buffer.WriteString("(")
		for i, child := range node.nodes {
			travel(&child, buffer, values)
			if i != n {
				switch node.Kind {
				case sql.AndKind:
					buffer.WriteString(" and ")
				case sql.OrKind:
					buffer.WriteString(" or ")
				}
			}
		}
		buffer.WriteString(")")
	}
}

func (node *Node) Sql() (string, []interface{}) {
	var builder strings.Builder
	var values []interface{}
	if node.Kind>>32 == 0 {
		stringOp(node, &builder, &values)
		return builder.String(), values
	}
	builder.Grow(100)
	travel(node, &builder, &values)
	return builder.String(), values
}

func (node *Node) Sql1() (string, []interface{}) {
	var builder = bytess.NewByteBuffer(make([]byte, 60))
	var stack = datastruct.NewStack()
	var values []interface{}
	if node.Kind>>32 == 0 {
		return builder.String(), values
	}
	type nodeStru struct {
		node *Node
		idx  int
	}
	builder.WriteString("(")
	fmt.Println(len(node.nodes))
	stack = stack.Push(nodeStru{node, 0})

	var v interface{}
	var flag = false
	var i int
	var ns nodeStru
walk:
	v, stack = stack.Pop()
	if v == nil {
		goto exit
	}

	ns = v.(nodeStru)
	fmt.Println(len(ns.node.nodes))
	for i = ns.idx; i < len(ns.node.nodes); i++ {
		n := ns.node.nodes[i]
		if n.Kind>>32 == 0 {
			if i != 0 {
				switch n.Kind {
				case sql.AndKind:
					builder.WriteString(" and ")
				case sql.OrKind:
					builder.WriteString(" or ")
				}
			}
			stringOp1(&n, builder)
		} else {
			fmt.Println(n.Key)
			flag = true
			builder.WriteString("(")
			stack = stack.Push(nodeStru{&n, i + 1})
			stack = stack.Push(nodeStru{&n.nodes[i], 0})
			goto walk
		}
	}
	builder.WriteString(")")
	helper.Pass(flag)
exit:
	return builder.String(), values
}
