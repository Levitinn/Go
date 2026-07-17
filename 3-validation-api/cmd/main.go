package main

import (
	"3-validation-api/config"
	"3-validation-api/internal/verify"
	"3-validation-api/storage"
	"fmt"
	"net/http"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		fmt.Printf("failed to create config: %v", err)
		return
	}
	store := storage.NewStorage("./storage/tokens.json")
	handler := verify.NewHandler(verify.HandlerDeps{Config: config, Storage: store})

	mux := http.NewServeMux()
	mux.HandleFunc("POST /send", handler.Send)
	mux.HandleFunc("GET /verify/{hash}", handler.Verify)
	fmt.Println("server started on port 8081")
	http.ListenAndServe(":8081", mux)
}
