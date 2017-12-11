package database


import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
	"fmt"
)

func ConnectDB() *sql.DB {
	dbConnect := os.Getenv("DB_CONNECTION")
	address := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"),
		os.Getenv("SSLMODE"),
		os.Getenv("DB_PASSWORD"),
	)
	db, err := sql.Open(dbConnect, address)
	if err != nil {
		panic(err)
	}
	return db
}
