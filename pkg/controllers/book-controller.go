package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/matahariramadhan/go-book-management-api/pkg/models"
)

var Book *models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {

}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := models.FindAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBookById(w http.ResponseWriter, r *http.Request) {

}
func UpdateBook(w http.ResponseWriter, r *http.Request) {

}
func DeleteBook(w http.ResponseWriter, r *http.Request) {

}
