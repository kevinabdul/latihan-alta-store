package controllers

import (
	"net/http"
	"fmt"

	product "lataltastore/services/product"
	models "lataltastore/models"

	"github.com/labstack/echo/v4"
)

func GetProductsController(c echo.Context) error {
	categoryName := c.QueryParam("category")
	fmt.Println(categoryName)
	productsTarget, rowsAffected, err := product.GetProducts(categoryName)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if rowsAffected == int64(0) {
		return c.JSON(http.StatusBadRequest, struct {
			Status 	string
			Message string
		}{Status: "failed", Message: "No product found with the provided category. Please check your product category name"})
	}

	return c.JSON(http.StatusOK, struct {
		Status 		string
		Products 	[]models.ProductAPI
	}{Status: "success", Products: productsTarget})
}
