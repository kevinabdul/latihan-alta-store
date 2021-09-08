package models

import (
	"time"
)

type Product struct {
	ID			uint 		`gorm:"primaryKey"`
	ProductName	string		`gorm:"type:varchar(70);not null;unique" json:"product_name" form:"product_name"`
	Category	string		`gorm:"type:varchar(70);not null" json:"category" form:"category"`
	Price 		uint 		`gorm:"type:int unsigned;not null" json:"price" form:"price"`
	CreatedAt	time.Time 	
	UpdatedAt	time.Time	
}

type ProductAPI struct {
	ProductName	string		`gorm:"type:varchar(70);not null;unique" json:"product_name" form:"product_name"`
	Category	string		`gorm:"type:varchar(70);not null" json:"category" form:"category"`
	Price 		uint 		`gorm:"type:int unsigned;not null" json:"price" form:"price"`
}