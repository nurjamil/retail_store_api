package controllers

import (
	"net/http"
	"retailStore/config"
	"retailStore/middlewares"
	"retailStore/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetOrderController(c echo.Context) error {
	id := uint(middlewares.ExtractTokenUserId(c))
	order := models.Order{
		UserID: id,
	}
	c.Bind(&order)
	orders := []models.Order{}

	if err := config.DB.Preload("Courier").Preload("Address").Preload("PaymentService").Preload("OrderItem.Item.ItemCategory").Preload("Payment").Where(&order).Find(&orders).Error; err != nil {
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
		"message": "success deleting items",
		"data":    orders,
	})

}

func PostOrderController(c echo.Context) error {
	id := uint(middlewares.ExtractTokenUserId(c))
	order := models.Order{
		UserID: id,
	}
	c.Bind(&order)

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

	orderTemp := models.Order{}

	if err := config.DB.Preload("Courier").Preload("Address").Preload("PaymentService").Preload("OrderItem.Item.ItemCategory").Preload("Payment").Where("id = ?", order.ID).First(&orderTemp).Error; err != nil {
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

func DeleteOrderController(c echo.Context) error {
	id := uint(middlewares.ExtractTokenUserId(c))
	order := models.Order{
		UserID: id,
	}
	c.Bind(&order)
	if err := config.DB.Preload("Courier").Preload("Address").Preload("PaymentService").Preload("OrderItem.Item.ItemCategory").Preload("Payment").First(&order).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "failed",
			"message": "Order Not Found",
			"data":    "",
		})
	}

	if err := config.DB.Preload("Courier").Preload("Address").Preload("PaymentService").Preload("OrderItem.Item.ItemCategory").Preload("Payment").Unscoped().Delete(&order).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "failed",
			"message": "Order Not Found",
			"data":    "",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "success deleting orders",
		"data":    order,
	})

}
