package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	internal "github.com/boraxpr/go-web-service/internal/dao"
)

// @Summary Get all quotations
// @Description Returns all quotations
// @Accept  json
// @Produce  json
// @Success 200 {array} Quotation
// @Router /quotation [get]
func GetAllQuotations(quotationDAO internal.Dao[internal.Quotation]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return

		}
		fmt.Println("GetAllQuotationsHandler")

		result, err := quotationDAO.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonResult, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResult)
	}
}

// @Summary Get a quotation
// @Description Returns a quotation
// @Accept  json
// @Produce  json
// @Success 200 {object} Quotation
// @Router /quotation/{id} [get]
func GetQuotationById(quotationDAO internal.Dao[internal.Quotation]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		fmt.Println("GetQuotationHandler")
		id := strings.TrimPrefix(r.URL.Path, "/quotation/")

		if id == "" {
			http.Error(w, "Invalid quotation ID", http.StatusBadRequest)
			return
		}
		id32, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := quotationDAO.Get(uint32(id32))
		if err != nil {
			if err.Error() == "scanning one: no rows in result set" {
				http.Error(w, "Quotation not found", http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonResult, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResult)
	}
}
