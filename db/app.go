// db/app.go
package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// App struct holds the dependencies for the application
type App struct {
	DB *pgxpool.Pool
}
