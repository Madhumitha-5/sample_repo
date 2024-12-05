package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:Natural@5@tcp(127.0.0.1:3306)/sample_repo"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the MySQL database: %v", err)
	}
	log.Println("MySQL database connected successfully!")
}
