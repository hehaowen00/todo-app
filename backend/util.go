package main

import (
	"net/http"
)

func httpMessage(w http.ResponseWriter, statusCode int) {
	http.Error(w, http.StatusText(statusCode), statusCode)
}
