package controllers

import (
	"net/http"
	"retailStore/lib/db"
	"retailStore/middlewares"
	"retailStore/models"

	"github.com/labstack/echo"
)

func ResponUserSuccess(status uint, message string, data models.User) models.UserResponse {
	userResponse := models.UserResponse{
		Code:    status,
		Status:  "success",
		Message: message,
		Data:    data,
	}
	return userResponse
}

func GetUserDetailController(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	user, err := db.GetUserDetail(int(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   user,
	})
}

func CreateUserController(c echo.Context) error {
	user, err := db.CreateUser(c)

	// c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusCreated, ResponUserSuccess(http.StatusCreated, "user account created", user.(models.User)))
}

func UpdateUserDetailController(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	user, err := db.UpdateUserDetail(int(id), c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "user accoount updated",
		"data":    user,
	})
}

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	_, err := db.LoginUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, ResponUserSuccess(http.StatusOK, "sucess login", user))
}
