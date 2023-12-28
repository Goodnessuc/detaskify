package utils

import "net/http"

func JSONContentTypeWrapper(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set the Content-Type header to JSON
		w.Header().Set("Content-Type", "application/json")
		handler(w, r)
	}
}
