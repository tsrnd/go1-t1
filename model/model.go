package model

import (
	"database/sql"
)

var DB *sql.DB

func SetDatabase(database *sql.DB) {
	DB = database
}
