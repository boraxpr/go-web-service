package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	internal "github.com/boraxpr/go-web-service/internal/dao"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v5"
)

// @Summary Get all quotations
// @Description Returns all quotations
// @Accept  json
// @Produce  json
// @Success 200 {array} Quotation
// @Router /quotation [get]
func GetAllQuotations(quotationDAO internal.Dao[internal.Quotation]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		result, err := quotationDAO.GetAll()
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, &ErrResponse{Err: err})
			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, result)
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
		fmt.Println("GetQuotationHandler")
		idParam := chi.URLParam(r, "id")
		if idParam == "" {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, ErrInvalidRequest(nil))
			return
		}
		id32, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, ErrInvalidRequest(err))
			return
		}

		result, err := quotationDAO.Get(uint32(id32))
		if err != nil && err != pgx.ErrNoRows {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, &ErrResponse{Err: err})
			return
		}
		if err == pgx.ErrNoRows {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, ErrNotFound)
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, result)
	}
}
