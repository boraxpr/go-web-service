package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", defaultHandler) // Catch-all route
	http.HandleFunc("ping", pingHandler)
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "404 Not Found - Page not found :(" )
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}