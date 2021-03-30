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

	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserByIdController)

	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserByIdController)
	e.PUT("/users/:id", controllers.UpdateUserByIdController)

	e.GET("/items/:id", controllers.GetItemWIthParamsController)
	e.GET("/items", controllers.GetItemController)
	e.POST("/items", controllers.PostItemController)

	e.GET("/shoppingcarts", controllers.GetShoppingCartController)
	//e.POST("/shoppingcarts", controllers.PostItemToShoppingCartController)

	e.GET("/couriers", controllers.GetCouriersController)
	e.GET("/couriers/:id", controllers.GetCourierByIdController)
	e.DELETE("/couriers/:id", controllers.DeleteCourierByIdController)
	e.PUT("/couriers/:id", controllers.UpdateCourierByIdController)
	e.POST("/couriers", controllers.CreateCourierController)

	e.GET("/itemCategories", controllers.GetItemCategoriesController)
	e.GET("/itemCategories/:id", controllers.GetItemCategoryByIdController)
	e.DELETE("/itemCategories/:id", controllers.DeleteItemCategoryByIdController)
	e.PUT("/itemCategories/:id", controllers.UpdateItemCategoryByIdController)
	e.POST("/itemCategories", controllers.CreateItemCategoryController)

	eJWT := e.Group("") 
	eJWT.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	return e
}
