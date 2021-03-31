package controllers

import (
	"net/http"
	"retailStore/config"
	"retailStore/lib/db"
	"retailStore/middlewares"
	"retailStore/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetItemWIthParamsController(c echo.Context) error {
	category := c.Param("id")
	val, _ := strconv.Atoi(category)
	model := models.Item{
		ItemCategoryID: uint(val),
	}
	items := []models.Item{}
	err := config.DB.Where(&model).Find(&items)
	if err.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "failed",
			"message": "failed getting items",
			"data":    "",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "success getting items",
		"data":    items,
	})

}

func GetItemController(c echo.Context) error {
	item := []models.Item{}

	if err := config.DB.Find(&item).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "getting items",
		"data":    item,
	})

}

func PostItemController(c echo.Context) error {
	role := middlewares.ExtractTokenUserRole(c)
	if role == "admin" {
		item, err := db.CreateItem(c)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"status":  "failed",
				"message": "bad request",
				"data":    "",
			})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"code":    http.StatusCreated,
			"status":  "success",
			"message": "item created",
			"data":    item,
		})
	
	}
	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
		"code":    http.StatusUnauthorized,
		"status":  "failed",
		"message": "you have no permission",
	})
}
