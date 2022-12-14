package routes

import (
	"github.com/Ibrahim-Muhammad13/books/pkg/controller"
	"github.com/gorilla/mux"
)

var RegisterBookRoute = func(router *mux.Router) {
	router.HandleFunc("/books", controller.GetAllBooks).Methods("GET")
	router.HandleFunc("/books", controller.CreateBook).Methods("POST")
	router.HandleFunc("/books/{bookId}", controller.GetBookById).Methods("GET")
	router.HandleFunc("/books/{bookId}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookId}", controller.DeleteBook).Methods("DELETE")
}
