package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type myData struct {
	Id   int
	Name string
}

func main() {
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db, err)
}
