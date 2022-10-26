package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matahariramadhan/go-book-management-api/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
