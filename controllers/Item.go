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
	id, _ := strconv.Atoi(c.Param("id"))
	model := models.Item{}

	if id > 0 {
		model.ID = uint(id)
	}
	items := []models.Item{}

	if err := config.DB.Preload("ItemCategory").Where(&model).Find(&items).Error; err != nil {
		return c.JSON(http.StatusBadRequest, models.ItemResponse{
			Code:    http.StatusBadRequest,
			Status:  "failed",
			Message: "failed getting items",
		})
	}

	itemResponseArr := models.ItemResponseArr{
		Code:    200,
		Status:  "success",
		Message: "success getting items",
		Data:    items,
	}
	return c.JSON(http.StatusOK, itemResponseArr)

}

func GetItemController(c echo.Context) error {
	categoryName := c.QueryParam("category")

	id, _ := strconv.Atoi(c.QueryParam("id"))
	model := models.Item{}
	if id > 0 {
		model.ID = uint(id)
	}

	minPrice, _ := strconv.Atoi(c.QueryParam("min_price"))
	maxPrice, _ := strconv.Atoi(c.QueryParam("max_price"))

	itemCategory := models.ItemCategory{
		CategoryName: categoryName,
	}
	if err := config.DB.Where(&itemCategory).First(&itemCategory).Error; err != nil {
		return c.JSON(http.StatusBadRequest, models.ItemResponse{
			Code:    http.StatusBadRequest,
			Status:  "failed",
			Message: "failed getting items, itemcategory not found",
		})
	}
	if categoryName != "" {
		model.ItemCategoryID = itemCategory.ID
	}

	items := []models.Item{}
	var err error
	if minPrice > 0 && maxPrice > 0 {
		err = config.DB.Preload("ItemCategory").Where(&model).Where("price >= ? AND price <= ?", minPrice, maxPrice).Find(&items).Error
	} else if minPrice > 0 {
		err = config.DB.Preload("ItemCategory").Where(&model).Where("price >= ?", minPrice).Find(&items).Error
	} else if maxPrice > 0 {
		err = config.DB.Preload("ItemCategory").Where(&model).Where("price <= ?", maxPrice).Find(&items).Error
	} else {
		err = config.DB.Preload("ItemCategory").Where(&model).Find(&items).Error
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponFailure("failed getting items"))
	}

	return c.JSON(http.StatusOK, models.ItemResponseArr{
		Code:    200,
		Status:  "success",
		Message: "success getting items",
		Data:    items,
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
		"data":    "",
	})
}
