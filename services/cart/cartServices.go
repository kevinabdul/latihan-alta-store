package services

import (
	"lataltastore/config"
	"lataltastore/models"
	"fmt"
)

func GetCart(userId int) (models.Cart, int64, error) {
	var cart models.Cart

	res := config.Db.Model(&models.Cart{}).Find(&cart, userId)

	if res.Error != nil {
		return models.Cart{}, res.RowsAffected, res.Error
	}
	if res.RowsAffected == 0 {
		return models.Cart{}, res.RowsAffected, res.Error
	}
	return cart, res.RowsAffected, nil
}

func CreateCart(userId uint) (uint, error) {
	newCart := models.Cart{}
	newCart.UserId = userId
	res := config.Db.Create(&newCart)
	
	if res.Error != nil {
		return 0, res.Error
	}
	
	return newCart.Id, nil
}