package utils

import (
	"encoding/json"
	"net/http"
)

// response writer for json
func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// middleware to handle errors and response
func MakeHttpHandlerFunc(f ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusInternalServerError, ApiError{Error: err.Error()})
		}
	}
}

type ApiError struct {
	Error string
}

type ApiFunc func(http.ResponseWriter, *http.Request) error
