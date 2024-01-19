package models

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Username     string `gorm:"not null"`
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string `gorm:"not null"`
	Role         string `gorm:"not null"`
}
