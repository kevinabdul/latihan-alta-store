package services

import (
	"lataltastore/config"
	"lataltastore/models"
)

func GetProducts() ([]models.ProductAPI, error) {
	var products []models.ProductAPI

	res := config.Db.Model(&models.Product{}).Find(&products)

	if res.Error != nil {
		return []models.ProductAPI{}, res.Error
	}

	return products, nil
}

func GetProductById(productId int) (models.ProductAPI, int64, error) {
	var product models.ProductAPI

	res := config.Db.Model(&models.Product{}).Where(`id = ?`, productId).Find(&product)

	if res.Error != nil {
		return models.ProductAPI{}, res.RowsAffected, res.Error
	}

	if res.RowsAffected == 0 {
		return models.ProductAPI{}, res.RowsAffected, res.Error
	}
	return product, res.RowsAffected, nil
}