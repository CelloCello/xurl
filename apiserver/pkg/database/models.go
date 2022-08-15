package database

import (
	"time"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	Url       string    `gorm:"size:100;not null;unique" json:"url"`
	Code      string    `gorm:"size:100;not null;" json:"code"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
