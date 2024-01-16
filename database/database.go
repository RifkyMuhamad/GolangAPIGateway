package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	dsn := os.Getenv("MYSQL_URL")
	db , err2 := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err2 != nil {
		panic(err2)
	}

	fmt.Println("Database Connection")

	return db
}

func CobaENV() {
	dsn := os.Getenv("MYSQL_LOCAL")
	fmt.Println(dsn)
}