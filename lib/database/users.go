package database

import (
	"errors"
	"github.com/labstack/echo/v4"
	"orm-crud/config"
	"orm-crud/middleware"
	"orm-crud/models"
	"strconv"
)

func GetUser() (*[]models.User, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return &[]models.User{}, err
	}
	return &users, nil
}

func GetUserById(c echo.Context) (*models.User, error) {
	var users models.User
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return &models.User{}, err
	}
	return &users, nil
}

func CreateUser(c echo.Context) (*models.User, error) {
	var users models.User
	c.Bind(&users)
	if err := config.DB.Save(&users).Error; err != nil {
		return &models.User{}, err
	}
	return &users, nil
}

func UpdateUser(c echo.Context) (*models.User, error) {
	var users models.User
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&users)
	if err := config.DB.Where("id = ?", id).Updates(&users).Error; err != nil {
		return &models.User{}, err
	}
	if err := config.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return &models.User{}, errors.New("Data tidak ditemukan")
	}
	return &users, nil
}

func DeleteUser(c echo.Context) (*models.User, error) {
	var users models.User
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return &models.User{}, errors.New("Data tidak ditemukan")
	} else if err = config.DB.Where("id = ?", id).Delete(&users).Error; err != nil {
		return &models.User{}, err
	}
	return &users, nil
}

func GetDetailUsers(userId int) (interface{}, error) {
	var user models.User
	if err := config.DB.Find(&user, userId).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func LoginUsers(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middleware.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err = config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
