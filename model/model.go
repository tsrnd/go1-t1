package model

import (
	"database/sql"
	"time"
)

type Model struct {
	ID        uint      `schema:"id"`
	CreatedAt time.Time `schema:"created_at"`
	UpdatedAt time.Time `schema:"updated_at"`
	DeletedAt time.Time `schema:"deleted_at"`
}

var DBCon *sql.DB

func SetDatabase(database *sql.DB) {
	DBCon = database
}
