package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	err := godotenv.Load(".env")
    if err != nil {
        panic("Error loading .env file")
    }

	dsn := os.Getenv("MYSQL_URL")
	fmt.Println(dsn)
	db , err2 := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err2 != nil {
		panic(err2)
	}

	fmt.Println("Database Connection")

	return db
}