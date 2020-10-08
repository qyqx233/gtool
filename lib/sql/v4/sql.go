package sql

import (
	"fastweb/lib/bytess"
	"fastweb/lib/sql/v1"
	"strings"
)

type Mul struct {
	Kind  int
	Key   string
	Value interface{}
	And   []Mul
	Or    []Mul
}

func Eq(key string) Mul {
	return Mul{Kind: sql.EqKind, Key: key}
}

func Neq(key string) Mul {
	return Mul{Kind: sql.NeqKind, Key: key}
}

func In(key string) Mul {
	return Mul{Kind: sql.InKind, Key: key}
}

func Nin(key string) Mul {
	return Mul{Kind: sql.NinKind, Key: key}
}

func stringOp1(o *Mul, buffer *bytess.ByteBuffer, vs *bytess.ValueBuffer) {
	buffer.WriteString(o.Key)
	vs.WriteValue(&o.Value)
}

func stringOp(o *Mul, buffer *bytess.ByteBuffer, vs *bytess.ValueBuffer) {
	switch o.Kind {
	case sql.EqKind:
		buffer.WriteString(o.Key)
		// buffer.WriteString(" = ?")
		if o.Value != nil {
			// vs.WriteValue(o.Value)
		}
	case sql.NeqKind:
		buffer.WriteString(o.Key)
		buffer.WriteString(" != ?")
		if o.Value != nil {
			vs.WriteValue(o.Value)
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

func (m *Mul) string(b *bytess.ByteBuffer, vs *bytess.ValueBuffer) {
	for i, v := range m.And {
		if i != 0 {
			b.WriteString(" and ")
		}
		if v.Key == "" {
			b.WriteString("(")
			v.string(b, vs)
			b.WriteString(")")
		} else {
			stringOp1(&v, b, vs)
			// b.WriteString(m.Key)
		}
	}
	for i, v := range m.Or {
		if i != 0 {
			b.WriteString(" or ")
		}
		if v.Key == "" {
			b.WriteString("(")
			v.string(b, vs)
			b.WriteString(")")
		} else {
			b.WriteString(m.Key)
		}
	}
}
