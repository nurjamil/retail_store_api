package routes

import (
	"os"
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

	e.GET("/items", controllers.GetItemController)
	e.GET("/items/:id", controllers.GetItemWIthParamsController)

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
	eJWT.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT"))))

	eJWT.POST("/address", controllers.CreateAddressController)
	eJWT.GET("/address", controllers.GetAddressController)
	eJWT.GET("/address/:id", controllers.GetAddressByIdController)

	eJWT.GET("/users", controllers.GetUserDetailController)
	eJWT.PUT("/users", controllers.UpdateUserDetailController)

	eJWT.GET("/shoppingcarts", controllers.GetShoppingCartController)
	eJWT.POST("/shoppingcarts", controllers.PostItemToShoppingCartController)
	eJWT.POST("/shoppingcarts/checkout", controllers.ShoppingCartCheckoutController)
	eJWT.DELETE("/shoppingcarts", controllers.DeleteItemFromShoppingCartController)

	eJWT.GET("/orders", controllers.GetOrderController)
	eJWT.POST("/orders", controllers.PostOrderController)
	eJWT.DELETE("/orders", controllers.DeleteOrderController)

	eJWT.GET("/payments", controllers.GetPaymentController)

	eAdmin := eJWT.Group("")
	eAdmin.POST("/items", controllers.PostItemController)

	return e
}
