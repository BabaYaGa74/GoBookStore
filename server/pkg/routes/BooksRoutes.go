package routes

import (
	"bookstore-go/server/pkg/handlers"

	"github.com/gorilla/mux"
)


var RegisterBookStoreRoutes = func(router *mux.Router){

	router.HandleFunc("/books", handlers.CreateBookHandler).Methods("POST");
	router.HandleFunc("/books", handlers.GetBooksHandler).Methods("GET");
	router.HandleFunc("/book/{bookId}", handlers.GetSingleBookHandler).Methods("GET");
	router.HandleFunc("/book/{bookId}", handlers.UpdateBookHandler).Methods("PUT");
	router.HandleFunc("/book/{bookId}", handlers.DeleteBookHandler).Methods("DELETE");
	
}