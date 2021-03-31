package routes

import (
	"retailStore/constants"
	"retailStore/controllers"
	"retailStore/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	middlewares.LogMiddlewares(e)

	e.POST("/register", controllers.CreateUserController)
	e.POST("/login", controllers.LoginUserController)

	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserByIdController)
	e.DELETE("/users/:id", controllers.DeleteUserByIdController)
	e.PUT("/users/:id", controllers.UpdateUserByIdController)

	e.GET("/items", controllers.GetItemController)
	e.GET("/items/:id", controllers.GetItemWIthParamsController)
	e.POST("/items", controllers.PostItemController) // admin

	e.GET("/couriers", controllers.GetCouriersController)
	e.GET("/couriers/:id", controllers.GetCourierByIdController)
	e.DELETE("/couriers/:id", controllers.DeleteCourierByIdController)
	e.PUT("/couriers/:id", controllers.UpdateCourierByIdController)
	e.POST("/couriers", controllers.CreateCourierController)

	eJWT := e.Group("")
	eJWT.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	eJWT.GET("/shoppingcarts", controllers.GetShoppingCartController) //authenticated user
	eJWT.POST("/shoppingcarts", controllers.PostItemToShoppingCartController)
	eJWT.DELETE("/shoppingcarts", controllers.DeleteItemFromShoppingCartController)

	eJWT.GET("/orders", controllers.GetOrderController)
	eJWT.POST("/orders", controllers.PostOrderController)
	eJWT.DELETE("/orders", controllers.DeleteOrderController)
	return e
}
