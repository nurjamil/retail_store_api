package db

import (
	"retailStore/config"
	"retailStore/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetCouriers() (interface{}, interface{}) {
	couriers := []models.Courier{}

	if err := config.DB.Find(&couriers).Error; err != nil {
		return nil, err
	}
	return couriers, nil
}
func GetCourierById(c echo.Context) (interface{}, interface{}) {
	id, _ := strconv.Atoi(c.Param("id"))
	courier := models.Courier{}

	if err := config.DB.Find(&courier, id).Error; err != nil {
		return nil, err
	}
	return courier, nil
}

func CreateCourier(c echo.Context) (interface{}, interface{}) {
	courier := models.Courier{}
	c.Bind(&courier)
	if err := config.DB.Create(&courier).Error; err != nil {
		return nil, err
	}
	
	return courier, nil
}

func DeleteCourierById(c echo.Context) (interface{}, interface{}) {
	id, _ := strconv.Atoi(c.Param("id"))
	courier := models.Courier{}

	if err := config.DB.Delete(&courier, id).Error; err != nil {
		return nil, err
	}
	return courier, nil
}

func UpdateCourierById(c echo.Context) (interface{}, interface{}) {
	id, _ := strconv.Atoi(c.Param("id"))
	courier := models.Courier{}
	c.Bind(&courier)
	courier.ID = uint(id)
	if err := config.DB.Save(&courier).Error; err != nil {
		return nil, err
	}
	return courier, nil
}

