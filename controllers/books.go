package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"orm-crud/lib/database"
)

func GetBookController(c echo.Context) error {
	books, err := database.GetBook()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed",
			"error":   err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get Books",
		"books":   books,
	})
}

func CreateBookController(c echo.Context) error {
	books, err := database.CreateBook(c)
	c.Bind(&books)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed",
			"error":   err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Create Book",
		"books":   books,
	})
}

func UpdateBookController(c echo.Context) error {
	books, err := database.UpdateBook(c)
	c.Bind(&books)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Update Book",
		"books":   books,
	})
}

func DeleteBookController(c echo.Context) error {
	books, err := database.DeleteBook(c)
	c.Bind(&books)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Delete Book",
		"books":   books,
	})
}
