package controllers

import (
	"net/http"
	"retailStore/config"
	"retailStore/middlewares"
	"retailStore/models"

	"github.com/labstack/echo"
)

func GetShoppingCartController(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	model := models.ShoppingCart{
		ID: uint(id),
	}

	shoppingCart := models.ShoppingCart{}
	err := config.DB.Preload("ShoppingCartList.Item.ItemCategory").Where(&model).First(&shoppingCart)
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

func getUserID() uint {
	return 18
}

func PostItemToShoppingCartController(c echo.Context) error {
	cartList := models.ShoppingCartList{}
	c.Bind(&cartList)
	cart := models.ShoppingCart{
		UserID: getUserID(),
	}
	if err := config.DB.Where(&cart).First(&cart).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "failed",
			"message": err.Error,
			"data":    "",
		})
	}
	cartList.ShoppingCartID = cart.ID
	if err := config.DB.Save(&cartList).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "failed",
			"message": err.Error,
			"data":    "",
		})
	}
	if err := config.DB.Preload("Item.ItemCategory").Where("shopping_cart_id = ?", cart.ID).Find(&cart.ShoppingCartList).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "failed",
			"message": "bambang",
			"data":    "",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "success getting items",
		"data":    cart,
	})
}
