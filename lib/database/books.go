package database

import (
	"github.com/labstack/echo/v4"
	"orm-crud/config"
	"orm-crud/models"
	"strconv"
)

func GetBook() (*[]models.Book, error) {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return &[]models.Book{}, err
	}
	return &books, nil
}

func GetBookById(c echo.Context) (*models.Book, error) {
	books := models.Book{}
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Where("id = ?", id).First(&books).Error; err != nil {
		return &models.Book{}, err
	}
	return &books, nil
}

func CreateBook(c echo.Context) (*models.Book, error) {
	books := models.Book{}
	c.Bind(&books)
	if err := config.DB.Save(&books).Error; err != nil {
		return &models.Book{}, err
	}
	return &books, nil
}

func UpdateBook(c echo.Context) (*models.Book, error) {
	var books models.Book
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&books)
	if err := config.DB.Where("id = ?", id).Updates(&books).Error; err != nil {
		return &models.Book{}, err
	}
	return &books, nil
}

func DeleteBook(c echo.Context) (*models.Book, error) {
	var books models.Book
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Where("id = ?", id).Delete(&books).Error; err != nil {
		return &models.Book{}, err
	}
	return &books, nil
}
