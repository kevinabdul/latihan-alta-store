package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Id		uint 		`gorm:"primaryKey`
	UserId	uint 		
}