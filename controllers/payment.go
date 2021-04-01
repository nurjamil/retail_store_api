package controllers

import (
	"net/http"
	"retailStore/config"
	"retailStore/middlewares"
	"retailStore/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetPaymentController(c echo.Context) error {
	id := uint(middlewares.ExtractTokenUserId(c))
	idPayment, _ := strconv.Atoi(c.QueryParam("payment_id"))
	status := c.QueryParam("status")
	payment := models.Payment{
		ID:     uint(idPayment),
		UserID: id,
		Status: status,
	}

	payments := []models.Payment{}

	if err := config.DB.Where(&payment).Find(&payments).Error; err != nil {
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
		"message": "success getting items",
		"data":    payments,
	})
}
