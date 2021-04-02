package controllers

import (
	"errors"
	"net/http"
	"retailStore/config"
	"retailStore/middlewares"
	"retailStore/models"
	"strconv"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func ResponFailure(message string) models.ShoppingCartAPI {
	model := models.ShoppingCartAPI{
		Code:    http.StatusBadRequest,
		Status:  "failed",
		Message: message,
	}
	return model
}

func ResponSuccess(res interface{}) models.ShoppingCartAPI {
	modelResponse := models.ShoppingCartAPI{
		Code:    200,
		Status:  "success",
		Message: "sucess getting shopping cart",
		Data:    res.(models.ShoppingCart),
	}
	return modelResponse
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
	newCartList := models.ShoppingCartList{}

	//check if item already on cartlist
	if err := config.DB.Where("shopping_cart_id = ? AND item_id = ?", cart.ID, cartList.ItemID).First(&newCartList).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusBadRequest, ResponFailure(err.Error()))
		}
	}
	if newCartList.ID != 0 {
		newCartList.Quantity += cartList.Quantity
		if err := config.DB.Save(&newCartList).Error; err != nil {
			return c.JSON(http.StatusBadRequest, ResponFailure(err.Error()))
		}
		if err := config.DB.Preload("Item.ItemCategory").Where("shopping_cart_id = ?", cart.ID).Find(&cart.ShoppingCartList).Error; err != nil {
			return c.JSON(http.StatusBadRequest, ResponFailure(err.Error()))
		}
		return c.JSON(http.StatusOK, ResponSuccess(cart))
	}

	//cartlist not found, save to database
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
			return c.JSON(http.StatusBadRequest, ResponFailure("failed deleting item from cart"))
		}
	}
	return c.JSON(http.StatusOK, models.ShoppingCartAPI{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success deleting items from cart",
	})
}

func ShoppingCartCheckoutController(c echo.Context) error {
	id := uint(middlewares.ExtractTokenUserId(c))
	order := models.Order{
		UserID: id,
	}
	c.Bind(&order)
	shoppingCart := models.ShoppingCart{
		UserID: id,
	}
	if err := config.DB.Preload("ShoppingCartList.Item.ItemCategory").Where(&shoppingCart).First(&shoppingCart).Error; err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseOrder{
			Code:    http.StatusBadRequest,
			Status:  "failed",
			Message: err.Error(),
		})
	}

	order.OrderItem = make([]models.OrderItem, len(shoppingCart.ShoppingCartList))
	for pos, each := range shoppingCart.ShoppingCartList {
		order.OrderItem[pos].ItemID = each.ItemID
		order.OrderItem[pos].Quantity = each.Quantity
	}

	//check price and stock
	var totalAmount uint = 0
	for _, oneItem := range order.OrderItem {
		item := models.Item{}
		if err := config.DB.Where("id = ?", oneItem.ItemID).First(&item).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"status":  "failed",
				"message": err,
				"data":    "",
			})
		}
		if item.Stock < oneItem.Quantity {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"status":  "failed",
				"message": item.Name + " stock kurang",
				"data":    "",
			})
		}
		totalAmount += item.Price * oneItem.Quantity
	}
	order.TotalAmount = totalAmount
	if err := config.DB.Create(&order).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "failed",
			"message": err,
			"data":    "",
		})
	}

	payment := models.Payment{
		UserID:          order.UserID,
		OrderID:         order.ID,
		TransactionCode: strconv.Itoa(int(id)) + strconv.Itoa(int(order.ID)) + strconv.Itoa(int(order.CourierID)),
		Status:          "Belum Dibayar",
		TotalAmount:     order.TotalAmount,
	}

	if err := config.DB.Create(&payment).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "failed",
			"message": err,
			"data":    "",
		})
	}

	orderTemp := models.Order{
		ID: order.ID,
	}

	if err := config.DB.Preload("Courier").Preload("Address").Preload("PaymentService").Preload("OrderItem.Item.ItemCategory").Preload("Payment").Where(&orderTemp).First(&orderTemp).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "failed",
			"message": err,
			"data":    "",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "success creating order",
		"data":    orderTemp,
	})

}
