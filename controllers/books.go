package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"orm-crud/lib/database"
	"strconv"
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
		"data":    newResponseArray(*books),
	})
}

func GetBookByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := database.GetBookById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Find Book ById",
		"data":    newResponse(*book),
	})
}

func CreateBookController(c echo.Context) error {
	var bookReq RequestBook
	c.Bind(&bookReq)
	result, err := database.CreateBook(bookReq.toModel())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed",
			"error":   err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Create Book",
		"data":    newResponse(*result),
	})
}

func UpdateBookController(c echo.Context) error {
	var bookReq RequestBook
	c.Bind(&bookReq)
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := database.UpdateBook(id, bookReq.toModel())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Update Book",
		"data":    newResponse(*result),
	})
}

func DeleteBookController(c echo.Context) error {
	var bookReq RequestBook
	c.Bind(&bookReq)
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := database.DeleteBook(id, bookReq.toModel())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Delete Book",
		/*"books":   books,*/
	})
}
