package controllers

import (
	"errors"
	"net/http"
	"retailStore/config"
	"retailStore/middlewares"
	"retailStore/models"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func ResponFailure(message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"status":  "failed",
		"message": message,
		"data":    "",
	}
}

func ResponSuccess(res interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "success getting items",
		"data":    res,
	}
}

func GetShoppingCartController(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	model := models.ShoppingCart{
		UserID: uint(id),
	}
	cart := models.ShoppingCart{}
	err := config.DB.Preload("ShoppingCartList.Item.ItemCategory").Where(&model).First(&cart)
	if err.Error != nil {
		return c.JSON(http.StatusBadRequest, ResponFailure(string(err.Error.Error())))
	}

	return c.JSON(http.StatusOK, ResponSuccess(cart))
}

func PostItemToShoppingCartController(c echo.Context) error {

	cartList := models.ShoppingCartList{}
	c.Bind(&cartList)

	id := middlewares.ExtractTokenUserId(c)
	cart := models.ShoppingCart{
		UserID: uint(id),
	}
	//take cart from user id and append it to model
	if err := config.DB.Where(&cart).First(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusBadRequest, ResponFailure("Record Not Found"))
		}
		return c.JSON(http.StatusBadRequest, ResponFailure(err.Error()))
	}

	cartList.ShoppingCartID = cart.ID
	newCartList := models.ShoppingCartList{
		ShoppingCartID: cart.ID,
	}

	//check if item already on cartlist
	if err := newCartList.Find(c, config.DB); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"status":  "failed",
				"message": err.Error,
				"data":    "",
			})
		}
	}
	if newCartList.ID != 0 {
		newCartList.Quantity += cartList.Quantity
		return c.JSON(http.StatusOK, ResponSuccess(newCartList))
	}

	//cartlist not found save to database
	if err := config.DB.Save(&cartList).Error; err != nil {
		return c.JSON(http.StatusBadRequest, ResponFailure(err.Error()))
	}

	//get all cartlist from cart
	if err := config.DB.Preload("Item.ItemCategory").Where("shopping_cart_id = ?", cart.ID).Find(&cart.ShoppingCartList).Error; err != nil {
		return c.JSON(http.StatusBadRequest, ResponFailure(err.Error()))
	}

	return c.JSON(http.StatusOK, ResponSuccess(cart))
}

func DeleteItemFromShoppingCartController(c echo.Context) error {
	id := uint(middlewares.ExtractTokenUserId(c))
	cart := models.ShoppingCart{
		UserID: id,
	}
	if err := config.DB.Where(&cart).First(&cart).Error; err != nil {
		return c.JSON(http.StatusBadRequest, ResponFailure(err.Error()))
	}
	items := []models.Item{}
	c.Bind(&items)
	for _, oneItem := range items {
		cartList := models.ShoppingCartList{}
		if err := config.DB.Where("shopping_cart_id = ? AND item_id = ?", cart.ID, oneItem.ID).Unscoped().Delete(&cartList).Error; err != nil {
			return c.JSON(http.StatusBadRequest, ResponFailure(err.Error()))
		}
	}
	return c.JSON(http.StatusOK, ResponSuccess(items))
}
