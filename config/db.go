package config

import (
	"database/sql"
	Db "goweb1/services/database/sql"

	_ "github.com/lib/pq"
)

// DB func
func DB() *sql.DB {
	// dbDlct := os.Getenv("DATABASE_DLCT")
	// dbUser := os.Getenv("DATABASE_USER")
	// dbPass := os.Getenv("DATABASE_PASS")
	// dbHost := os.Getenv("DATABASE_HOST")
	// dbPort := os.Getenv("DATABASE_PORT")
	// dbName := os.Getenv("DATABASE_NAME")
	//db, err := Db.Connect(dbDlct, dbUser, dbPass, dbHost, dbPort, dbName)
	//Connect(dlct, user, pass, name, host, port string)
	db := Db.Connect("postgres", "postgres", "123456", "goweb1", "127.0.0.1", "5432", "disable")
	return db
}
