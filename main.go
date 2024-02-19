package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/boraxpr/go-web-service/db"
	"github.com/boraxpr/go-web-service/handlers"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from a .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if PORT environment variable is not set
	}
	secret_key := os.Getenv("SECRET_KEY")
	// Connect to db
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_STRING"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	// Create an instance of App with the database connection
	app := &db.App{DB: conn}

	mux := http.NewServeMux()
	mux.Handle("/", handlers.LoggingMiddleware(http.HandlerFunc(handlers.DefaultHandler(app))))

	// Wrap the PingHandler with both the LoggingMiddleware and AuthMiddleware
	mux.Handle(
		"/ping",
		handlers.LoggingMiddleware(handlers.AuthMiddleware(http.HandlerFunc(handlers.PingHandler), secret_key)),
	)

	mux.Handle(
		"/quotation",
		handlers.LoggingMiddleware(
			handlers.AuthMiddleware(http.HandlerFunc(handlers.GetAllQuotationsHandler(app)), secret_key),
		),
	)

	fmt.Printf("Server listening on %s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
