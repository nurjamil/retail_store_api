package db

import (
	"retailStore/config"
	"retailStore/models"

	"github.com/labstack/echo"
)

func CreateItem(c echo.Context) (interface{}, interface{}) {
	item := models.Item{}
	c.Bind(&item)
	if err := config.DB.Save(&item).Error; err != nil {
		return nil, err
	}
	config.DB.Joins("ItemCategory").Where(&item).First(&item)

	return item, nil
}

