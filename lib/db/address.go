package db

import (
	"retailStore/config"
	"retailStore/models"
	"strconv"

	"github.com/labstack/echo"
)

func CreateAddress(id int, c echo.Context) (interface{}, error) {
	address := models.Address{
		UserID: uint(id),
	}
	c.Bind(&address)
	if err := config.DB.Save(&address).Error; err != nil {
		return nil, err
	}
	return address, nil
}

func GetAddresses(id int) (interface{},error){
	addresses := []models.Address{}

	if err := config.DB.Where("user_id = ?", id).Find(&addresses).Error; err != nil {
		return nil,err
	}
	return addresses, nil
}

func GetAddressById(id int, c echo.Context) (interface{}, error){
	paramId, _ := strconv.Atoi(c.Param("id"))
	address := models.Address{}

	if err := config.DB.Where("user_id = ? AND id = ?", id, paramId).First(&address).Error; err != nil {
		return nil, err
	}
	return address, nil
}