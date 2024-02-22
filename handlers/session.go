package handlers

import (
	"net/http"
)

func SessionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SessionHandler"))
}
