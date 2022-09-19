package database

import (
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Initialize(dbFile string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Link{})
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}

func TestDBInit() *gorm.DB {
	db := Initialize("test.db")
	// sqlDB, err := db.DB()
	// if err != nil {
	// 	panic("failed to get test db")
	// }
	// sqlDB.SetMaxIdleConns(3)
	// sqlDB.LogMode(true)
	return db
}

func TestDBFree(db *gorm.DB) error {
	// db.Close()
	err := os.Remove("test.db")
	return err
}
