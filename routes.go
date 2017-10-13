package main

import (
	"fmt"
	"net/http"
)

func startRoutes() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hola mundo")
}
