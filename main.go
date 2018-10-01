package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func main() {

	db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	a := App{}
	a.Initialize()

	a.Run(":8080")
}