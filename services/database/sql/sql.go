package sql

import (
	"database/sql"
	"fmt"
)

// Connect func
func Connect(dlct, user, pass, name, host, port, sslmode string) *sql.DB {
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s password=%s",
		host, user, name, port, sslmode, pass)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}
