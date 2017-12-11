package migrations

import (
	_ "github.com/lib/pq"
	"fmt"
)

func CreateOrdersTable() {
	stmt, err := DB.Prepare(`CREATE TABLE orders(
		id SERIAL PRIMARY KEY,
		total integer NOT NULL,
		status boolean NOT NULL,
		user_id integer,
		created_at timestamp with time zone,
		updated_at timestamp with time zone,
		deleted_at timestamp with time zone
	);`)
	checkErr(err)
	stmt.Exec()
	fmt.Println("Created table orders")
}