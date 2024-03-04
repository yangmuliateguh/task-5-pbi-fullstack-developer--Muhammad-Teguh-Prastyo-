package models

import (
	"time"

	"gorm.io/gorm"
)

// Photo represents the photo model
type Photo struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `gorm:"foreignKey:UserID;" json:"-"`
}
