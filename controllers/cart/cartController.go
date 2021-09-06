package controllers

import (
	"net/http"
	"strconv"

	cart "lataltastore/services/cart"
	models "lataltastore/models"

	"github.com/labstack/echo/v4"
)

func GetCartController(c echo.Context) error {
	userId , _ := strconv.Atoi(c.Request().Header.Get("userId"))
	cartTarget, rowsAffected, err := cart.GetCart(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if rowsAffected == int64(0) {
		return c.JSON(http.StatusBadRequest, struct {
			Status 	string
			Message 	string
		}{Status: "failed", Message: "wrong user id"})
	}

	return c.JSON(http.StatusOK, struct {
		Status 	string
		Cart 	models.CartAPI
	}{Status: "success", Cart: cartTarget})
}

func CreateCartController(c echo.Context) error {
	userId , _ := strconv.Atoi(c.Request().Header.Get("userId"))
	id, err := cart.CreateCart(uint(userId))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, struct {
		Status 	string
		Message string
		CartId 	uint
	}{Status: "success", Message: "Cart has been created!", CartId: id})

}