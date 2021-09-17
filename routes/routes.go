package routes

import (
	"github.com/labstack/echo/v4"
	"orm-crud/controllers"
)

func NewRouters() *echo.Echo {
	e := echo.New()
	e.GET("/book", controllers.GetBookController)
	e.GET("/book/:id", controllers.GetBookByIdController)
	e.POST("/book", controllers.CreateBookController)
	e.PUT("/book", controllers.UpdateBookController)
	e.DELETE("/book", controllers.DeleteBookController)
	return e
}
