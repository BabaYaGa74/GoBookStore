package main

import (
	"bookstore-go/server/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
	})
	handler := c.Handler(r)
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", handler)
	log.Fatal(http.ListenAndServe("localhost:5000", handler))
}
