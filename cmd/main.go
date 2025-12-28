package main

import (
	"log"
	"net/http"

	_ "github.com/Archiker-715/rest-api/docs"
	"github.com/Archiker-715/rest-api/internal/handlers"
	"github.com/Archiker-715/rest-api/internal/middleware"
	"github.com/Archiker-715/rest-api/internal/repository/pg"
	"github.com/Archiker-715/rest-api/internal/repository/pg/items"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// for swagger
// swag init --generalInfo ./cmd/main.go

// @title My API
// @version 1.0
// @description Описание API
// @host localhost:8080
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	pg.Connect()

	itemRepo := items.NewItemRepository(pg.DB)
	itemHandler := handlers.NewItemHandler(itemRepo)

	r := mux.NewRouter()

	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.CORSMiddleware)
	r.Use(middleware.SecurityMiddleware)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/items", itemHandler.GetItems).Methods("GET")
	api.HandleFunc("/items", itemHandler.CreateItem).Methods("POST")
	api.HandleFunc("/items/{id}", itemHandler.UpdateItem).Methods("PUT")
	api.HandleFunc("/items/{id}", itemHandler.DeleteItem).Methods("DELETE")

	log.Println("Server starting :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
