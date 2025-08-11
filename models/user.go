package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	gorm.Model
	Name     string    `json:"name" gorm:"size:100;not null"`
	Email    string    `json:"email" gorm:"size:100;not null;uniqueIndex"`
	Address  string    `json:"address" gorm:"size:255"`
	JoinedAt time.Time `json:"joined_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
}
