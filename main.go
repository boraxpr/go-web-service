package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	_ "github.com/boraxpr/go-web-service/docs"
	"github.com/boraxpr/go-web-service/handlers"
	dao "github.com/boraxpr/go-web-service/internal/dao"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title Go Web Service API
// @version 1.0
// @description This is a sample server for a Go web service.

// @host
// @BasePath /
func main() {

	port := os.Getenv("PORT")
	intPort, err := strconv.Atoi(port)
	if err != nil {
		intPort = 8080 // Default to port 8080 if PORT environment variable is not set
	}

	// Connect to db
	pool, err := pgxpool.New(context.Background(), os.Getenv("DB_STRING"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()
	// // Create an instance of App with the database connection
	// app := &db.App{DB: pool}

	quotationDAO := dao.NewQuotationDao(pool)
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))
	// r.Use(render.SetContentType(render.ContentTypeJSON))
	// Public Routes
	r.Group(func(r chi.Router) {
		r.Get("/swagger*", handlers.SwaggerAuth(httpSwagger.Handler()))
		// r.Handle("/", handlers.LoggingMiddleware(http.HandlerFunc(handlers.Default(app))))
		r.Post(
			"/login", handlers.LoginHandler(),
		)

	})

	// Private Routes
	// Require authentication
	r.Group(func(r chi.Router) {
		r.Use(handlers.AuthMiddleware)
		r.Get(
			"/quotation", handlers.GetAllQuotations(quotationDAO),
		)
		r.Get(
			"/quotation/{id}", handlers.GetQuotationById(quotationDAO))
		r.Post(
			"/session", handlers.SessionHandler,
		)
	})

	// Apply CORS middleware to your router
	handler := corsMiddleware(r)
	fmt.Printf("Server listening on %s\n", strconv.Itoa(intPort))
	if err := http.ListenAndServe(":"+ strconv.Itoa(intPort), handler); err != nil {
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
