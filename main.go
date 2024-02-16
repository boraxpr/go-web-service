package main

import (
	"fmt"
	"net/http"

	"github.com/boraxpr/go-web-service/handlers"
)

func main() {
	mux := http.NewServeMux()

	// Wrap the DefaultHandler with the LoggingMiddleware
	mux.Handle("/", handlers.LoggingMiddleware(http.HandlerFunc(handlers.DefaultHandler)))

	// Wrap the PingHandler with both the LoggingMiddleware and AuthMiddleware
	mux.Handle("/ping", handlers.LoggingMiddleware(handlers.AuthMiddleware(http.HandlerFunc(handlers.PingHandler))))

	fmt.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
