package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB = nil

func InitDB() {
	db_instance, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db = db_instance
}

func GetDB() *gorm.DB {
	if db != nil {
		return db
	}
	InitDB()
	return db
}
