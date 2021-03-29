package controllers

import (
	"net/http"
	"retailStore/config"
	"retailStore/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetShoppingCartController(c echo.Context) error {
	id := c.QueryParam("shopping_cart_id")
	val, _ := strconv.Atoi(id)
	model := models.ShoppingCart{
		ID: uint(val),
	}

	shoppingCart := models.ShoppingCart{}
	err := config.DB.Where(&model).First(&shoppingCart)
	if err.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "failed",
			"message": err.Error,
			"data":    "",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "success getting items",
		"data":    shoppingCart,
	})

}
