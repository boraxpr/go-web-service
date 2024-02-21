package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/boraxpr/go-web-service/db"
)

type Quotation struct {
	DocNum       uint32          `db:"doc_num"`
	CreatedDate  sql.NullTime    `db:"created_date"`
	Status       sql.NullString  `db:"status"`
	Currency     sql.NullString  `db:"currency"`
	ProjectName  sql.NullInt64   `db:"project_name"`
	GrandTotal   sql.NullFloat64 `db:"grand_total"`
	CustomerID   sql.NullString  `db:"customer_id"`
	DueDate      sql.NullTime    `db:"due_date"`
	UpdatedAt    sql.NullTime    `db:"updated_at"`
	IsActive     sql.NullBool    `db:"is_active"`
	CreditDay    sql.NullInt64   `db:"credit_day"`
	Remark       sql.NullString  `db:"remark"`
	Note         sql.NullString  `db:"note"`
	Attachment   sql.NullString  `db:"attachment"`
	UpdatedBy    sql.NullString  `db:"updated_by"`
	Running      sql.NullString  `db:"running"`
	CustomerName sql.NullString  `db:"customer_name"`
}

// @Summary Get all quotations
// @Description Returns all quotations
// @Accept  json
// @Produce  json
// @Success 200 {array} Quotation
// @Router /quotation [get]
func GetAllQuotationsHandler(app *db.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return

		}
		fmt.Println("GetAllQuotationsHandler")
		// Use db to query the database
		rows, err := app.DB.Query(
			context.Background(),
			"SELECT quotation.doc_num, quotation.created_date, quotation.status, quotation.currency, quotation.project_name, quotation.grand_total, quotation.customer_id, quotation.due_date, quotation.updated_at, quotation.is_active, quotation.credit_day, quotation.remark, quotation.note, quotation.attachment, quotation.updated_by, quotation.running, customers.customer_name FROM quotation JOIN customers ON quotation.customer_id = customers.id",
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var result []Quotation
		for rows.Next() {
			var q Quotation
			err := rows.Scan(
				&q.DocNum,
				&q.CreatedDate,
				&q.Status,
				&q.Currency,
				&q.ProjectName,
				&q.GrandTotal,
				&q.CustomerID,
				&q.DueDate,
				&q.UpdatedAt,
				&q.IsActive,
				&q.CreditDay,
				&q.Remark,
				&q.Note,
				&q.Attachment,
				&q.UpdatedBy,
				&q.Running,
				&q.CustomerName,
			)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			result = append(result, q)
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
