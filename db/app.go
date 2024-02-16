// db/app.go
package db

import "github.com/jackc/pgx/v5"

// App struct holds the dependencies for the application
type App struct {
	DB *pgx.Conn
}
