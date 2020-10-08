package main

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/qyqx233/gtool/model"
)

func parseStruct(val interface{}) map[string]reflect.StructTag {
	object := reflect.ValueOf(val)
	ref := object.Elem()
	typeOfType := ref.Type()
	leng := ref.NumField()
	mp := make(map[string]reflect.StructTag, leng)
	for i := 0; i < leng; i++ {
		ft := typeOfType.Field(i)
		tag := ft.Tag
		fmt.Println(tag, ft.Name)
		// fv := ref.Field(i)
		// mp[key] = fv.Interface()
		// key := tag.Get("orm")
		mp[ft.Name] = tag

	}
	return mp
}

func packSql(v interface{}, table string) string {
	valMap := parseStruct(v)
	var bs = make([]byte, 0, 200)
	buf := bytes.NewBuffer(bs)
	for _, v := range valMap {
		buf.Write([]byte(v.Get("orm") + ", "))
	}
	// buf.Write([]byte("from " + table))
	s := buf.String()
	s = s[:len(s)-2]
	sql := "select " + s + " from " + table
	return sql
}

func TestS(t *testing.T) {
	var m model.Answer
	sql := packSql(&m, "a")
	t.Log(sql)
}

func TestDb(t *testing.T) {
	var answer model.Answer
	model.InitMyDB()
	err := model.MyDB.QueryRow("select id from question where id = ?", 1).Scan(&answer.Id)
	if err != nil {
		t.Error(err)
	}
	t.Log(answer.Id)
}
