package routes

import (
	"github.com/gorilla/mux"
	"github.com/matahariramadhan/go-book-management-api/pkg/controllers"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
