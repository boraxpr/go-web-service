package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/boraxpr/go-web-service/db"
)

type Quotation struct {
	DocNum      int64           `db:"doc_num"`
	CreatedDate sql.NullTime    `db:"created_date"`
	Status      sql.NullString  `db:"status"`
	Currency    sql.NullString  `db:"currency"`
	ProjectName sql.NullInt64   `db:"project_name"`
	GrandTotal  sql.NullFloat64 `db:"grand_total"`
	CustomerID  sql.NullString  `db:"customer_id"`
	DueDate     sql.NullTime    `db:"due_date"`
	UpdatedAt   sql.NullTime    `db:"updated_at"`
	IsActive    sql.NullBool    `db:"is_active"`
	CreditDay   sql.NullInt64   `db:"credit_day"`
	Remark      sql.NullString  `db:"remark"`
	Note        sql.NullString  `db:"note"`
	Attachment  sql.NullString  `db:"attachment"`
	UpdatedBy   sql.NullString  `db:"updated_by"`
	Running     sql.NullString  `db:"running"`
}

func QuotationHandler(app *db.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return

		}
		// Use db to query the database
		rows, err := app.DB.Query(context.Background(), "SELECT * FROM quotation")
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
