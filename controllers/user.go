package controllers

import (
	"fmt"
	"net/http"
	"retailStore/lib/db"

	"github.com/labstack/echo"
)



func GetUsersController(c echo.Context) error {
	users, err := db.GetUsers()

	if err != nil {
		return c.JSON(http.StatusBadRequest,map[string]interface{}{
			"status":"failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":"success",
		"data":users,
	})
}
func GetUserByIdController(c echo.Context) error {
	user, err := db.GetUserById(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest,map[string]interface{}{
			"status":"failed",
			"message": "bad request",
		})
	}
	fmt.Println(user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":"success",
		"data":user,
	})
}


func CreateUserController(c echo.Context) error {
	user, err := db.CreateUser(c)

	// c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest,map[string]interface{}{
			"status":"failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message":"user accoount created",
		"data":user,
	})
}

func DeleteUserByIdController(c echo.Context) error {
	user, err := db.DeleteUserById(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest,map[string]interface{}{
			"status":"failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":"success",
		"data":user,
	})
}


func UpdateUserByIdController(c echo.Context) error {
	user, err := db.UpdateUserById(c)

	// c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest,map[string]interface{}{
			"status":"failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message":"user accoount updated",
		"data":user,
	})
}

