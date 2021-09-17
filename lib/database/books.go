package database

import (
	"github.com/labstack/echo/v4"
	"orm-crud/config"
	"orm-crud/models"
)

func GetBook() (*[]models.Book, error) {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return &[]models.Book{}, err
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
	c.Bind(&books)
	if err := config.DB.Where("id = ?", &books.ID).Updates(&books).Error; err != nil {
		return &models.Book{}, err
	}
	return &books, nil
}

func DeleteBook(c echo.Context) (*models.Book, error) {
	var books models.Book
	c.Bind(&books)
	if err := config.DB.Where("id = ?", &books.ID).Delete(&books).Error; err != nil {
		return &models.Book{}, err
	}
	return &books, nil
}
