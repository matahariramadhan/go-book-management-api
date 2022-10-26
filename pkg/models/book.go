package models

import (
	"github.com/matahariramadhan/go-book-management-api/pkg/config"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func FindAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}
