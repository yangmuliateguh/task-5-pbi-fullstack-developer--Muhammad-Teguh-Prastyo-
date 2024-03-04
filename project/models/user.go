package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents the user model
type User struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `json:"username" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Photos    []Photo   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"photos"`
}
