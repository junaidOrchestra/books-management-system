package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func NewSQLiteConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to SQLite:", err)
	}

	return db
}
