package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID               string `gorm:"primarykey;type:uuid;default:gen_random_uuid()"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	Username         string
	PasswordHash     string
	Email            string
	BookingHistories []BookingHistory
}
