package db

import (
	"retailStore/config"
	"retailStore/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetUsers() (interface{}, interface{}) {
	users := []models.User{}

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func GetUserById(c echo.Context) (interface{}, interface{}) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}

	if err := config.DB.Find(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(c echo.Context) (interface{}, interface{}) {
	user := models.User{}
	c.Bind(&user)
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	user.ShoppingCart.UserID = user.ID

	if err := config.DB.Create(&user.ShoppingCart).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUserById(c echo.Context) (interface{}, interface{}) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}

	if err := config.DB.Delete(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUserById(c echo.Context) (interface{}, interface{}) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}
	c.Bind(&user)
	user.ID = uint(id)
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
