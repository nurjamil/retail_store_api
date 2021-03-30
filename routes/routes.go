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

	e.GET("/shoppingcarts", controllers.GetShoppingCartController) //authenticated user
	//e.POST("/shoppingcarts", controllers.PostItemToShoppingCartController)

	eJWT := e.Group("") 
	eJWT.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	return e
}
