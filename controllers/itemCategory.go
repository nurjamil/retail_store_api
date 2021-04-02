package controllers

import (
	"net/http"
	"retailStore/lib/db"
	"retailStore/middlewares"

	"github.com/labstack/echo"
)

func GetItemCategoriesController(c echo.Context) error {
	role := middlewares.ExtractTokenUserRole(c)
	if role == "admin" {
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
	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
		"code":    http.StatusUnauthorized,
		"status":  "failed",
		"message": "you have no permission",
		"data":    "",
	})
}


func GetItemCategoryByIdController(c echo.Context) error {
	role := middlewares.ExtractTokenUserRole(c)
	if role == "admin" {
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
	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
		"code":    http.StatusUnauthorized,
		"status":  "failed",
		"message": "you have no permission",
		"data":    "",
	})
}

func CreateItemCategoryController(c echo.Context) error {
	role := middlewares.ExtractTokenUserRole(c)
	if role == "admin" {
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
	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
		"code":    http.StatusUnauthorized,
		"status":  "failed",
		"message": "you have no permission",
		"data":    "",
	})
}

func DeleteItemCategoryByIdController(c echo.Context) error {
	role := middlewares.ExtractTokenUserRole(c)
	if role == "admin" {
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
	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
		"code":    http.StatusUnauthorized,
		"status":  "failed",
		"message": "you have no permission",
		"data":    "",
	})
}

func UpdateItemCategoryByIdController(c echo.Context) error {
	role := middlewares.ExtractTokenUserRole(c)
	if role == "admin" {
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
	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
		"code":    http.StatusUnauthorized,
		"status":  "failed",
		"message": "you have no permission",
		"data":    "",
	})
}

