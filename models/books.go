package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID     uint   `gorm:"primaryKey"`
	Title  string `json:"title"`
	Isbn   string `json:"isbn"`
	Writer string `json:"writer"`
}
