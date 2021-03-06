package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"orm-crud/models"
)

var DB *gorm.DB

const JWT_SECRET string = "secretofjwt"
const JWT_EXP int = 1

func InitDB() {
	connectionString := "root:@tcp(127.0.0.1:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitalMigration() {
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.User{})
}
