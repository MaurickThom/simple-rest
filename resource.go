package main

import (
	"net/http"
)

// process handler for all request
func process(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		cs := getAll()

		w.Header().Set("Content-type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(cs.String()))
	}
}
