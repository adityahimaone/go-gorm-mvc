package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Age      int    `json:"age"`
	Sex      string `json:"sex"`
	//client_id int
}
