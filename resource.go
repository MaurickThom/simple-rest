package main

import (
	"net/http"
)

// process handler for all request
func process(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c := Contact{
			ID:    1,
			Name:  "LEIDY",
			City:  "BOGOTA",
			Phone: "4444444",
		}

		w.Header().Set("Content-type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(c.String()))
	}
}
