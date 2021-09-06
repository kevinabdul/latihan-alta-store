package services

import (
	"lataltastore/config"
	"lataltastore/models"
)

func GetCart(userId int) (models.CartAPI, int64, error) {
	var cart models.CartAPI

	res := config.Db.Model(&models.Cart{}).Where(`user_id = ?`, userId).Find(&cart)

	if res.Error != nil {
		return models.CartAPI{}, res.RowsAffected, res.Error
	}
	if res.RowsAffected == 0 {
		return models.CartAPI{}, res.RowsAffected, res.Error
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