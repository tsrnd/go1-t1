package migrations

import (
	_ "github.com/lib/pq"
	"fmt"
)

func CreateCategoriesTable() {
	stmt, err := DB.Prepare(`CREATE TABLE categories (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		created_at timestamp with time zone,
		updated_at timestamp with time zone,
		deleted_at timestamp with time zone
	);`)
	checkErr(err)
	stmt.Exec()
	fmt.Println("Created table categories")
}