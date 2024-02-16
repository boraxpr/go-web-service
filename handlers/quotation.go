package handlers

import (
	"fmt"
	"net/http"
)

func QuotationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")

}
