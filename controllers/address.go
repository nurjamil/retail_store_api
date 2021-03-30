package controllers

import (
	"net/http"
	"retailStore/lib/db"
	"retailStore/middlewares"

	"github.com/labstack/echo"
)

func CreateAddressController(c echo.Context) error {
	id:= middlewares.ExtractTokenUserId(c)
	address, err := db.CreateAddress(int(id), c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "failed",
			"message": "bad request",
			"data":    "",
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    http.StatusCreated,
		"status":  "success",
		"message": "address created",
		"data":    address,
	})
}


func GetAddressController(c echo.Context) error {
	id:= middlewares.ExtractTokenUserId(c)
	addresses,err := db.GetAddresses(int(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "getting addresses",
		"data":    addresses,
	})

}

func GetAddressByIdController(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	user, err := db.GetAddressById(int(id),c)

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