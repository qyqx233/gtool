package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var MyDB *sql.DB

func InitMyDB() {
	var err error
	MyDB, err = sql.Open("mysql", "root:123456@tcp(localhost:3307)/apps?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
}
