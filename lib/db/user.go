package db

import (
	"retailStore/config"
	"retailStore/middlewares"
	"retailStore/models"

	"github.com/labstack/echo"
)

func GetUserDetail(id int) (interface{}, error) {
	user := models.User{}

	if err := config.DB.Find(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(c echo.Context) (interface{}, error) {
	user := models.User{
		Role: "user",
	}
	c.Bind(&user)
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	user.ShoppingCart.UserID = user.ID

	if err := config.DB.Create(&user.ShoppingCart).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUserDetail(id int,c echo.Context) (interface{}, error) {
	user := models.User{}
	c.Bind(&user)
	user.ID = uint(id)
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}


func LoginUser(user *models.User) (interface{}, error){
	var err error
	if err = config.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(user).Error; err != nil {
		return nil, err
	}
	user.Token, err = middlewares.CreateToken(int(user.ID), user.Role)
	if err != nil {
		return nil,err
	}
	if err := config.DB.Save(user).Error; err!=nil{
		return nil, err
	}
	return user,nil
}