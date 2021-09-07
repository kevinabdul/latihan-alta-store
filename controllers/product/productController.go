package controllers

import (
	"net/http"
	"strconv"

	product "lataltastore/services/product"
	models "lataltastore/models"

	"github.com/labstack/echo/v4"
)

func GetProductsController(c echo.Context) error {
	res, err := product.GetProducts()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, struct {
		Status 		string
		Products 	[]models.ProductAPI
	}{Status: "success", Products: res})
}

func GetProductByIdController(c echo.Context) error {
	productId , _ := strconv.Atoi(c.Param("id"))
	productTarget, rowsAffected, err := product.GetProductById(productId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if rowsAffected == int64(0) {
		return c.JSON(http.StatusBadRequest, struct {
			Status 	string
			Message string
		}{Status: "failed", Message: "No product found with the provided id. Please check your product id"})
	}

	return c.JSON(http.StatusOK, struct {
		Status 	string
		Product models.ProductAPI
	}{Status: "success", Product: productTarget})
}
