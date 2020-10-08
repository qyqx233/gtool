package main

import (
	"fastweb/model"
)

//go:generate echo 100
func main() {
	model.InitMyDB()
}
