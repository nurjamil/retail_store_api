package db

import (
	"retailStore/config"
	"retailStore/models"

	"strconv"

	"github.com/labstack/echo"
)

func GetItemCategoires() (interface{}, interface{}) {
	itemCategories := []models.ItemCategory{}

	if err := config.DB.Find(&itemCategories).Error; err != nil {
		return nil, err
	}
	return itemCategories, nil
}

func GetItemCategoryById(c echo.Context) (interface{}, interface{}) {
	id, _ := strconv.Atoi(c.Param("id"))
	itemcategory := models.ItemCategory{}

	if err := config.DB.Find(&itemcategory, id).Error; err != nil {
		return nil, err
	}
	return itemcategory, nil
}

func CreateItemCategory(c echo.Context) (interface{}, interface{}) {
	itemcategory := models.ItemCategory{}
	c.Bind(&itemcategory)
	if err := config.DB.Create(&itemcategory).Error; err != nil {
		return nil, err
	}
	
	return itemcategory, nil
}

func DeleteItemCategoryById(c echo.Context) (interface{}, interface{}) {
	id, _ := strconv.Atoi(c.Param("id"))
	itemcategory := models.ItemCategory{}

	if err := config.DB.Delete(&itemcategory, id).Error; err != nil {
		return nil, err
	}
	return itemcategory, nil
}

func UpdateItemCategoryById(c echo.Context) (interface{}, interface{}) {
	id, _ := strconv.Atoi(c.Param("id"))
	itemcategory := models.ItemCategory{}
	c.Bind(&itemcategory)
	itemcategory.ID = uint(id)
	if err := config.DB.Save(&itemcategory).Error; err != nil {
		return nil, err
	}
	return itemcategory, nil
}
