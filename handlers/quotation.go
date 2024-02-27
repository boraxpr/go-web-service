package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

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
