package models

import (
	"time"
)

type Product struct {
	ID			uint 		`gorm:"primaryKey"`
	Name		string		`gorm:"type:varchar(70);not null;unique" json:"name" form:"name"`
	Category	string		`gorm:"type:varchar(70);not null" json:"category" form:"category"`
	Price 		uint 		`gorm:"type:int unsigned;not null" json:"price" form:"price"`
	CreatedAt	time.Time 	
	UpdatedAt	time.Time	
}

type ProductAPI struct {
	Name		string		`gorm:"type:varchar(70);not null;unique" json:"name" form:"name"`
	Category	string		`gorm:"type:varchar(70);not null" json:"category" form:"category"`
	Price 		uint 		`gorm:"type:int unsigned;not null" json:"price" form:"price"`
}