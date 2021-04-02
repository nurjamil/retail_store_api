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

	eJWT := e.Group("")
	eJWT.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT"))))

	eJWT.POST("/address", controllers.CreateAddressController)
	eJWT.GET("/address", controllers.GetAddressController)
	eJWT.GET("/address/:id", controllers.GetAddressByIdController)

	couriers := eJWT.Group("/couriers")
	couriers.GET("", controllers.GetCouriersController)
	couriers.GET("/:id", controllers.GetCourierByIdController)
	couriers.DELETE("/:id", controllers.DeleteCourierByIdController)
	couriers.PUT("/:id", controllers.UpdateCourierByIdController)
	couriers.POST("", controllers.CreateCourierController)

	itemcategories := eJWT.Group("/itemcategories")
	itemcategories.GET("", controllers.GetItemCategoriesController)
	itemcategories.GET("/:id", controllers.GetItemCategoryByIdController)
	itemcategories.DELETE("/:id", controllers.DeleteItemCategoryByIdController)
	itemcategories.PUT("/:id", controllers.UpdateItemCategoryByIdController)
	itemcategories.POST("", controllers.CreateItemCategoryController)	
	
	eJWT.GET("/users", controllers.GetUserDetailController)
	
	eJWT.GET("/orders", controllers.GetOrderController)
	eJWT.POST("/orders", controllers.PostOrderController)
	eJWT.DELETE("/orders", controllers.DeleteOrderController)
	
	eJWT.GET("/payments", controllers.GetPaymentController)

	eJWT.GET("/shoppingcarts", controllers.GetShoppingCartController)
	eJWT.POST("/shoppingcarts", controllers.PostItemToShoppingCartController)
	eJWT.POST("/shoppingcarts/checkout", controllers.ShoppingCartCheckoutController)
	eJWT.DELETE("/shoppingcarts", controllers.DeleteItemFromShoppingCartController)

	eJWT.PUT("/users", controllers.UpdateUserDetailController)

	eAdmin := eJWT.Group("")
	eAdmin.POST("/items", controllers.PostItemController)

	return e
}
