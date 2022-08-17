package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Link struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid" json:"id"`
	Url       string    `gorm:"size:256;not null;unique" json:"url"`
	Code      string    `gorm:"size:50;not null;" json:"code"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

func (link *Link) BeforeCreate(tx *gorm.DB) (err error) {
	link.ID = uuid.New()
	return
}
