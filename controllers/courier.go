package controllers

import (
	"net/http"
	"retailStore/lib/db"

	"github.com/labstack/echo"
)

func GetCouriersController(c echo.Context) error {
	couriers, err := db.GetCouriers()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   couriers,
	})
}
func GetCourierByIdController(c echo.Context) error {
	courier, err := db.GetCourierById(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   courier,
	})
}

func CreateCourierController(c echo.Context) error {
	courier, err := db.CreateCourier(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "courier created",
		"data":    courier,
	})
}

func DeleteCourierByIdController(c echo.Context) error {
	courier, err := db.DeleteCourierById(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   courier,
	})
}

func UpdateCourierByIdController(c echo.Context) error {
	courier, err := db.UpdateCourierById(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "courier updated",
		"data":    courier,
	})
}

