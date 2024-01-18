package models

import (
	"time"
)

type Users struct {
	UserID      string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"` // UUID as a string
	FullName    string    `gorm:"type:varchar(100);not null"`
	Username    string    `gorm:"type:varchar(50);unique;not null"`
	Email       string    `gorm:"type:varchar(100);unique;not null"`
	Password    string    `gorm:"type:varchar(100);not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"` // Use this if you want to override the gorm.Model's CreatedAt
	UpdatedAt   time.Time `gorm:"autoUpdateTime"` // Use this if you want to override the gorm.Model's UpdatedAt
	ImageURL    string    `gorm:"type:text"`
	Points      int       `gorm:"default:0"`
	PhoneNumber string    `gorm:"type:varchar(20);"`
	Age         int       `gorm:"type:int"`
}

type UserData struct {
	UserID         string `gorm:"type:varchar(100);"`
	EducationLevel string `gorm:"type:varchar(100);"`
	WorkExp        string `gorm:"type:varchar(100);"`
	PrefIndustries string `gorm:"type:varchar(100);;"`
	TechSkills     string `gorm:"type:varchar(100);"`
	SoftSkills     string `gorm:"type:varchar(100)"`
	CareerGoals    string `gorm:"type:varchar(100)"`
	LimDesc        string `gorm:"type:varchar(100)"`
	PartBody       int    `gorm:"type:int"`
	PrefComm       string `gorm:"type:varchar(100);"`
	AccomNeeded    string `gorm:"type:varchar(100)"`
}
