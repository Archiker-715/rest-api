package main

import (
	"log"
	"net/http"

	"github.com/Archiker-715/rest-api/internal/handlers"
	"github.com/Archiker-715/rest-api/internal/middleware"
	"github.com/Archiker-715/rest-api/internal/repository/pg"
	"github.com/Archiker-715/rest-api/internal/repository/pg/items"
	"github.com/gorilla/mux"
)

func main() {
	pg.Connect()

	itemRepo := items.NewItemRepository(pg.DB)
	itemHandler := handlers.NewItemHandler(itemRepo)

	r := mux.NewRouter()

	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.CORSMiddleware)
	r.Use(middleware.SecurityMiddleware)

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/items", itemHandler.GetItems).Methods("GET")
	api.HandleFunc("/items", itemHandler.CreateItem).Methods("POST")
	api.HandleFunc("/items/{id}", itemHandler.UpdateItem).Methods("PUT")
	api.HandleFunc("/items/{id}", itemHandler.DeleteItem).Methods("DELETE")

	log.Println("Server starting :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
