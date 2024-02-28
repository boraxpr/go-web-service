package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/boraxpr/go-web-service/db"
	_ "github.com/boraxpr/go-web-service/docs"
	"github.com/boraxpr/go-web-service/handlers"
	dao "github.com/boraxpr/go-web-service/internal/dao"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title Go Web Service API
// @version 1.0
// @description This is a sample server for a Go web service.

// @host
// @BasePath /
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
	pool, err := pgxpool.New(context.Background(), os.Getenv("DB_STRING"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()
	// Create an instance of App with the database connection
	app := &db.App{DB: pool}

	quotationDAO := dao.NewQuotationDao(app)

	mux := http.NewServeMux()

	mux.Handle("/swagger/", handlers.SwaggerAuth(httpSwagger.Handler(), secret_key))
	mux.Handle("/", handlers.LoggingMiddleware(http.HandlerFunc(handlers.Default(app))))

	// Wrap the PingHandler with both the LoggingMiddleware and AuthMiddleware
	mux.Handle(
		"/login",
		handlers.LoggingMiddleware(
			http.HandlerFunc(handlers.LoginHandler(secret_key)),
		),
	)

	mux.Handle(
		"/quotation",
		handlers.LoggingMiddleware(
			handlers.AuthMiddleware(http.HandlerFunc(handlers.GetAllQuotations(quotationDAO)), secret_key),
		),
	)
	mux.Handle(
		"/quotation/",
		handlers.LoggingMiddleware(
			handlers.AuthMiddleware(http.HandlerFunc(handlers.GetQuotationById(quotationDAO)), secret_key),
		),
	)
	mux.Handle(
		"/session",
		handlers.LoggingMiddleware(
			handlers.AuthMiddleware(http.HandlerFunc(handlers.SessionHandler), secret_key),
		),
	)
	// Apply CORS middleware to your router
	handler := corsMiddleware(mux)
	fmt.Printf("Server listening on %s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().
			Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}
