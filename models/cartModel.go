package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Id		uint 		`gorm:"primaryKey`
	UserId	uint 		`gorm:"unique"`
}

type CartAPI struct {
	Id		uint 		`gorm:"primaryKey`
	UserId	uint 		`gorm:"unique"`
}