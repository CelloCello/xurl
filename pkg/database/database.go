package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	SqlSession *gorm.DB
)

func Initialize(dbFile string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
