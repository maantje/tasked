package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB

func ORM() *gorm.DB {
	return database
}

func Init() {
	db, err := gorm.Open(mysql.Open(os.Getenv("DB_DSN")), &gorm.Config{})

	database = db

	if err != nil {
		panic("Failed to establish database connection")
	}
}

func InitTest() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	database = db

	if err != nil {
		panic("Failed to establish database connection")
	}
}
