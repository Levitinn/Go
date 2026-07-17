package responses

import (
	"encoding/json"
	"net/http"
)

type SendResponse struct {
	Message string `json:"message"`
}

func JsonResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
