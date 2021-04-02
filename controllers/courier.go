package controllers

import (
	"net/http"
	"retailStore/lib/db"
	"retailStore/middlewares"

	"github.com/labstack/echo"
)

func GetCouriersController(c echo.Context) error {
	role := middlewares.ExtractTokenUserRole(c)
	if role == "admin" {
		couriers, err := db.GetCouriers()

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status":  "failed",
				"message": "bad request",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"status":  "success",
			"message": "success getting couriers",
			"data":   couriers,
		})
	}
	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
		"code":    http.StatusUnauthorized,
		"status":  "failed",
		"message": "you have no permission",
		"data":    "",
	})
	
}
func GetCourierByIdController(c echo.Context) error {
	role := middlewares.ExtractTokenUserRole(c)
	if role == "admin" {
		courier, err := db.GetCourierById(c)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status":  "failed",
				"message": "bad request",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"status":  "success",
			"message": "success getting courier by id",
			"data":   courier,
		})
	}
	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
		"code":    http.StatusUnauthorized,
		"status":  "failed",
		"message": "you have no permission",
		"data":    "",
	})
}

func CreateCourierController(c echo.Context) error {
	role := middlewares.ExtractTokenUserRole(c)
	if role == "admin" {
		courier, err := db.CreateCourier(c)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status":  "failed",
				"message": "bad request",
			})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"code":    200,
			"status":  "success",
			"message": "success courier created",
			"data":    courier,
		})
	}
	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
		"code":    http.StatusUnauthorized,
		"status":  "failed",
		"message": "you have no permission",
		"data":    "",
	})
}

func DeleteCourierByIdController(c echo.Context) error {
	role := middlewares.ExtractTokenUserRole(c)
	if role == "admin" {
		courier, err := db.DeleteCourierById(c)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status":  "failed",
				"message": "bad request",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"status":  "success",
			"message": "success deleting courier by id",
			"data":   courier,
		})
	}
	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
		"code":    http.StatusUnauthorized,
		"status":  "failed",
		"message": "you have no permission",
		"data":    "",
	})
}

func UpdateCourierByIdController(c echo.Context) error {
	role := middlewares.ExtractTokenUserRole(c)
	if role == "admin" {
		courier, err := db.UpdateCourierById(c)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status":  "failed",
				"message": "bad request",
			})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"code":    200,
			"status":  "success",
			"message": "success deleting courier by id",
			"data":    courier,
		})
	}
	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
		"code":    http.StatusUnauthorized,
		"status":  "failed",
		"message": "you have no permission",
		"data":    "",
	})
}

