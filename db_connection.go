package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var err error
var urlDSN = "root@tcp(localhost:3306)/go?parseTime=true"

func DataMigration() {
	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		panic("Error connecting to database")
	}
	Database.AutoMigrate(&Employee{})
	fmt.Println("connection established...")
}
