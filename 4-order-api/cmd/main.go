package main

import (
	"4-order-api/config"
	"4-order-api/internal/product"
	"4-order-api/pkg/db"
	"4-order-api/pkg/middleware"
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
	repo := product.NewRepository(db.DB)
	handler := product.NewHandler(product.HandlerDeps{Repository: repo})
	mux := http.NewServeMux()
	mux.HandleFunc("POST /products", handler.CreateProduct)
	mux.HandleFunc("PUT /products/{id}", handler.UpdateProduct)
	mux.HandleFunc("GET /products", handler.GetAllProducts)
	mux.HandleFunc("GET /products/{id}", handler.GetProductByID)
	mux.HandleFunc("DELETE /products/{id}", handler.DeleteProduct)
	log.Println("Server is running on port 8081")
	http.ListenAndServe(":8081", middleware.Logging(mux))
}
