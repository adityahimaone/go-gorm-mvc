package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"orm-crud/lib/database"
	"orm-crud/models"
	"strconv"
)

func GetAllUserController(c echo.Context) error {
	users, err := database.GetUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get All Users",
		"users":   users,
	})
}

func GetUserByIdController(c echo.Context) error {
	users, err := database.GetUserById(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Find Users ById",
		"user":    users,
	})
}

func CreateUserController(c echo.Context) error {
	users, err := database.CreateUser(c)
	c.Bind(&users)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Create Book",
		"user":    users,
	})
}

func UpdateUserController(c echo.Context) error {
	users, err := database.UpdateUser(c)
	c.Bind(&users)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Update Book",
		"user":    users,
	})
}

func DeleteUserController(c echo.Context) error {
	users, err := database.DeleteUser(c)
	c.Bind(&users)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Delete Book",
		/*"user":    users,*/
	})
}

func LoginUserController(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	users, err := database.LoginUsers(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Login",
		"user":    users,
	})
}

func GetUserDetailController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	users, err := database.GetDetailUsers(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"user":    users,
	})
}
