package main

import (
	"encoding/json"
	"net/http"
)

// Json is a helper function to send JSON responses
// w is the http.ResponseWriter
// r is the http.Request
// status is the HTTP status code
// payload is the JSON payload
func Json(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
