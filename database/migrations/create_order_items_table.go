package migrations

import (
	_ "github.com/lib/pq"
	"fmt"
)

func CreateOrderItemsTable() {
	stmt, err := DB.Prepare(`CREATE TABLE order_items (
		id SERIAL PRIMARY KEY,
		quantity integer NOT NULL,
		price integer NOT NULL,
		product_id integer NOT NULL,
		order_id integer NOT NULL,
		created_at timestamp with time zone,
		updated_at timestamp with time zone,
		deleted_at timestamp with time zone
	);`)
	checkErr(err)
	stmt.Exec()
	fmt.Println("Created table order items")
}