package routes

import (
	"github.com/gorilla/mux"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", createBook).Methods("POST")
	router.HandleFunc("/book", getAllBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", getBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", updateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", deleteBook).Methods("DELETE")
}
