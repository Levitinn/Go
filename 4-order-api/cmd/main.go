package main

import (
	"4-order-api/config"
	"4-order-api/internal/product"
	"4-order-api/pkg/db"
	"log"
	"net/http"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	db, err := db.NewDb(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if err := db.AutoMigrate(&product.Product{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("DB connected and migrated successfully")
	handler := product.NewHandler(product.HandlerDeps{DB: db.DB})
	mux := http.NewServeMux()
	mux.HandleFunc("POST /products", handler.CreateProduct)
	log.Println("Server is running on port 8081")
	http.ListenAndServe(":8081", mux)
}
