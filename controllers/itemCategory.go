package controllers

import (
	"net/http"
	"retailStore/lib/db"

	"github.com/labstack/echo"
)

func GetItemCategoriesController(c echo.Context) error {
	itemCategories, err := db.GetItemCategoires()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data": itemCategories,
	})
}


func GetItemCategoryByIdController(c echo.Context) error {
	itemcategory, err := db.GetItemCategoryById(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   itemcategory,
	})
}

func CreateItemCategoryController(c echo.Context) error {
	itemcategory, err := db.CreateItemCategory(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "itemcategory created",
		"data":    itemcategory,
	})
}

func DeleteItemCategoryByIdController(c echo.Context) error {
	itemcategory, err := db.DeleteItemCategoryById(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   itemcategory,
	})
}

func UpdateItemCategoryByIdController(c echo.Context) error {
	itemcategory, err := db.UpdateItemCategoryById(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "itemcategory updated",
		"data":    itemcategory,
	})
}

