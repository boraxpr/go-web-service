package internal

import (
	"context"
	"database/sql"

	"github.com/boraxpr/go-web-service/db"
	"github.com/jackc/pgx/v5/pgxpool"
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

type QuotationDao struct {
	db *pgxpool.Pool
}

func NewQuotationDao(app *db.App) Dao[Quotation] {
	return &QuotationDao{db: app.DB}
}

func (q *QuotationDao) Get(id uint32) (*Quotation, error) {
	// Implement logic to retrieve a quotation by its ID
	// Example: q.db.QueryRow("SELECT ... WHERE doc_num = $1", id).Scan(...)
	return nil, nil
}

func (q *QuotationDao) GetAll() ([]*Quotation, error) {
	// Implement logic to retrieve all quotations
	rows, err := q.db.Query(
		context.Background(),
		"SELECT quotation.doc_num, quotation.created_date, quotation.status, quotation.currency, quotation.project_name, quotation.grand_total, quotation.customer_id, quotation.due_date, quotation.updated_at, quotation.is_active, quotation.credit_day, quotation.remark, quotation.note, quotation.attachment, quotation.updated_by, quotation.running, customers.customer_name FROM quotation JOIN customers ON quotation.customer_id = customers.id",
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*Quotation
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
			return nil, err
		}
		result = append(result, &q)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (q *QuotationDao) Save(quotation *Quotation) error {
	// Implement logic to save a new quotation
	// Example: q.db.Exec("INSERT INTO ...", ...)
	return nil
}

func (q *QuotationDao) Update(quotation *Quotation) error {
	// Implement logic to update an existing quotation
	// Example: q.db.Exec("UPDATE ... SET ... WHERE doc_num = $1", quotation.DocNum)
	return nil
}

func (q *QuotationDao) Delete(id uint32) error {
	// Implement logic to delete a quotation by its ID
	// Example: q.db.Exec("DELETE FROM ... WHERE doc_num = $1", id)
	return nil
}
