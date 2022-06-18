package routes

import (
	"encoding/json"
	"net/http"
)

func jsonMessage(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(message)
}
