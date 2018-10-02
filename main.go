package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	a := App{}
	a.Initialize()

	a.Run()
	defer db.Close()
}