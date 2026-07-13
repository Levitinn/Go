package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func writeRandomNumber(w http.ResponseWriter, r *http.Request) {
	randomNumber := rand.Intn(6) + 1
	w.Write([]byte(fmt.Sprintf("%d", randomNumber)))
	return
}
func main() {
	http.HandleFunc("/random", writeRandomNumber)
	fmt.Println("Server is running on port 8081")
	http.ListenAndServe(":8081", nil)
}
