package services

import (
	"lataltastore/config"
	"lataltastore/models"
)

// func GetCarts() ([]models.CartAPI, error) {
// 	var carts []models.CartAPI

// 	res := config.Db.Model(&models.Cart{}).Find(&carts)

// 	if res.Error != nil {
// 		return []models.CartAPI{}, res.Error
// 	}

// 	return carts, nil
// }

func GetCartByUserId(userId int) ([]models.CartAPI, error) {
	var cart []models.CartAPI

	res := config.Db.Model(&models.Cart{}).Where(`user_id = ?`, userId).Find(&cart)

	if res.Error != nil {
		return []models.CartAPI{}, res.Error
	}

	return cart, nil
}

func UpdateCartByUserId(userCart []models.Cart, userId int)  (int64, error) {
	var rowsAffected int64

	for _, v := range userCart {
		if v.Quantity == 0 {
			continue
		}
		v.UserID = uint(userId)
		cart := models.Cart{}
		
		res := config.Db.Where(`user_id = ? AND product_name = ?`, userId, v.ProductName).Find(&cart)

		if res.Error != nil {
			return res.RowsAffected, res.Error
		}

		if res.RowsAffected == 0 {
			res2 := config.Db.Select("UserID", "ProductName", "Quantity").Create(&v)

			if res2.Error != nil {
				return res.RowsAffected, res2.Error
			}

			if res2.RowsAffected == 0 {
				return res.RowsAffected, res2.Error
			}
			rowsAffected++
		}

		if res.RowsAffected != 0 {
			res2 := config.Db.Model(&cart).Select("quantity").Updates(v)
			
			if res2.Error != nil {
				return res.RowsAffected, res.Error
			}

			if res2.RowsAffected == 0 {
				return res.RowsAffected, res.Error
			}

			rowsAffected++
		}
	}

	return rowsAffected, nil
}

// func CreateCart(userId uint) (uint, error) {
// 	newCart := models.Cart{}
// 	newCart.UserID = userId
// 	res := config.Db.Create(&newCart)
	
// 	if res.Error != nil {
// 		return 0, res.Error
// 	}
	
// 	return newCart.ID, nil
// }