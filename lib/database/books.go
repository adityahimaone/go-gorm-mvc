package database

import (
	"errors"
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

func GetBookById(id int) (*models.Book, error) {
	var book models.Book
	if err := config.DB.Where("id = ?", id).First(&book).Error; err != nil {
		return &models.Book{}, err
	}
	return &book, nil
}

func CreateBook(book models.Book) (*models.Book, error) {
	if err := config.DB.Save(&book).Error; err != nil {
		return &models.Book{}, err
	}
	return &book, nil
}

func UpdateBook(id int, book models.Book) (*models.Book, error) {
	if err := config.DB.Where("id = ?", id).Updates(&book).Error; err != nil {
		return &models.Book{}, err
	}

	if err := config.DB.Where("id = ?", id).First(&book).Error; err != nil {
		return nil, errors.New("Data tidak ditemukan")
	}
	return &book, nil
}

func DeleteBook(id int, book models.Book) (*models.Book, error) {
	if err := config.DB.Where("id = ?", id).First(&book).Error; err != nil {
		return &models.Book{}, errors.New("Data tidak ditemukan")
	}
	if err := config.DB.Where("id = ?", id).Delete(&book).Error; err != nil {
		return &models.Book{}, err
	}
	return &book, nil
}
