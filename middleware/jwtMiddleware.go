package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"orm-crud/config"
	"time"
)

func CreateToken(userID int ) (token string, err error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userid"] = userID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(config.JWT_EXP)).Unix()

	initToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = initToken.SignedString([]byte(config.JWT_SECRET))
	return
}

func ExtractTokenUserId(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(int)
		return userId
	}
	return 0
}
