package utils

import (
	"encoding/json"
	"net/http"
)

// Responder con un mensaje en formato JSON
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// Manejo de errores genéricos
func HandleError(w http.ResponseWriter, err error, message string, code int) {
	if err != nil {
		http.Error(w, message, code)
	}
}
