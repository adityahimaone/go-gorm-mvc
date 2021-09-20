package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"orm-crud/config"
	"orm-crud/controllers"
)

func NewRouters() *echo.Echo {
	e := echo.New()

	//Books Route
	e.GET("/book", controllers.GetBookController)
	e.GET("/book/:id", controllers.GetBookByIdController)
	e.POST("/book", controllers.CreateBookController)
	e.PUT("/book/:id", controllers.UpdateBookController)
	e.DELETE("/book/:id", controllers.DeleteBookController)

	//User Route
	e.GET("/user", controllers.GetAllUserController)
	e.GET("/user/:id", controllers.GetUserByIdController)
	e.POST("/user", controllers.CreateUserController)
	e.PUT("/user/:id", controllers.UpdateUserController)
	e.DELETE("/user/:id", controllers.DeleteUserController)

	//Login
	e.POST("/login", controllers.LoginUserController)
	//JWT AUth
	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(config.JWT_SECRET)))
	r.GET("/user/:id", controllers.GetUserDetailController)
	return e
}
