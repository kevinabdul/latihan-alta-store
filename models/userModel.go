package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id     		uint 	`json:"-" gorm:"primaryKey`
	Name   		string	`gorm:"type:varchar(50)" json:"name" form:"name"`
	Email 		string	`gorm:"unique;type:varchar(50);not null" json:"email" form:"email"`
	Password 	string	`gorm:"type:varchar(30);not null" json:"password" form: "password"`
	Cart		Cart
}

type UserAPI struct {
	Id 			uint 	`json:"id"`
	Name 		string 	`json:"name"`
	Email 		string 	`json:"email"`
}