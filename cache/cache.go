package cache

import "github.com/qyqx233/gtool/lib/convert"

type Cacher struct {
	*Cache
}

type Marshaler interface {
	MarshalMsg([]byte) ([]byte, error)
	UnmarshalMsg([]byte) ([]byte, error)
	Msgsize() int
}

func (c *Cacher) GetInt(fx func(Marshaler, uint64) error) func(Marshaler, uint64) error {
	return func(v Marshaler, s uint64) error {
		var key = convert.Uint642Bytes(s)
		var rs []byte
		var err error
		if rs = c.Get(rs, key); len(rs) != 0 {
			_, err = v.UnmarshalMsg(rs)
			return err
		}
		err = fx(v, s)
		if err != nil {
			return err
		}
		rs, err = v.MarshalMsg(rs)
		if err != nil {
			return err
		}
		c.Set(key, rs)
		return err
	}
}

func (c *Cacher) GetString(fx func(Marshaler, string) error) func(Marshaler, string) error {
	return func(v Marshaler, s string) error {
		var key = convert.String2Bytes(s)
		var rs []byte
		var err error
		if rs = c.Get(rs, key); len(rs) != 0 {
			_, err = v.UnmarshalMsg(rs)
			return err
		}
		err = fx(v, s)
		if err != nil {
			return err
		}
		rs, err = v.MarshalMsg(rs)
		if err != nil {
			return err
		}
		c.Set(key, rs)
		return err
	}
}
