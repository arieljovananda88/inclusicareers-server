package models

import (
	"time"
)

type User struct {
	UserID      string    `gorm:"type:uuid;primaryKey"` // UUID as a string
	FullName    string    `gorm:"type:varchar(100);not null"`
	Username    string    `gorm:"type:varchar(50);unique;not null"`
	Email       string    `gorm:"type:varchar(100);unique;not null"`
	Password    string    `gorm:"type:varchar(100);not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"` // Use this if you want to override the gorm.Model's CreatedAt
	UpdatedAt   time.Time `gorm:"autoUpdateTime"` // Use this if you want to override the gorm.Model's UpdatedAt
	ImageURL    string    `gorm:"type:text"`
	Points      int       `gorm:"default:0"`
	PhoneNumber string    `gorm:"type:varchar(20);unique"`
	Age         int       `gorm:"type:int"`
}
