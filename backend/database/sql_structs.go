package database

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type Stock struct {
	gorm.Model
	Symbol string `gorm:"unique;not null"`
}
