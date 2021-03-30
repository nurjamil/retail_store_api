package routes

import (
	"retailStore/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserByIdController)

	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserByIdController)
	e.PUT("/users/:id", controllers.UpdateUserByIdController)

	e.GET("/items/:id", controllers.GetItemWIthParamsController)
	e.GET("/items", controllers.GetItemController)
	e.POST("/items", controllers.PostItemController)

	e.GET("/shoppingcarts", controllers.GetShoppingCartController)
	e.POST("/shoppingcarts", controllers.PostItemToShoppingCartController)

	return e
}
