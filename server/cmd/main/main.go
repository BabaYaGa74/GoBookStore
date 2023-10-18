package main

import (
	"bookstore-go/server/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/cors"
)

func main(){
	r:= mux.NewRouter();	
	handler:= cors.Default().Handler(r);
	routes.RegisterBookStoreRoutes(r);
	http.Handle("/", handler);
	log.Fatal(http.ListenAndServe("localhost:5000", handler));
}