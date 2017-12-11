package migrations

import (
	_ "github.com/lib/pq"
	"fmt"
)

func CreateProductsTable() {
	stmt, err := DB.Prepare(`CREATE TABLE products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		price integer NOT NULL,
		image VARCHAR(255) NOT NULL,
		size integer NOT NULL,
		color VARCHAR(255) NOT NULL,
		category_id integer NOT NULL,
		created_at timestamp with time zone,
		updated_at timestamp with time zone,
		deleted_at timestamp with time zone
	);`)
	checkErr(err)
	stmt.Exec()
	fmt.Println("Created table products")
}