package handlers

import (
	"fmt"
	"net/http"

	"github.com/boraxpr/go-web-service/db"
)

func DefaultHandler(app *db.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "404 Not Found - Page not found")
	}
}
