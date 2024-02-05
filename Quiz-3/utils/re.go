// utils.go

package utils

import (
	"encoding/json"
	"net/http"
)

func SendJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func SendErrorResponse(w http.ResponseWriter, status int, message string) {
	SendJSONResponse(w, status, map[string]string{"error": message})
}
