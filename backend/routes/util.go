package routes

import (
	"encoding/json"
	"net/http"
)

func httpMessage(w http.ResponseWriter, statusCode int) {
	http.Error(w, http.StatusText(statusCode), statusCode)
}

func jsonMessage(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(message)
}
