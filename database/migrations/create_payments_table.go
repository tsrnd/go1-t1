package migrations

import (
	_ "github.com/lib/pq"
	"fmt"
)

func CreatePaymentsTable() {
	stmt, err := DB.Prepare(`CREATE TABLE payments (
		id SERIAL PRIMARY KEY,	
		method VARCHAR(255) NOT NULL,
		status integer NOT NULL,
		address VARCHAR(255) NOT NULL,
		total integer NOT NULL,
		vat integer NOT NULL,
		order_id integer NOT NULL,
		created_at timestamp with time zone,
		updated_at timestamp with time zone,
		deleted_at timestamp with time zone
	);`)
	checkErr(err)
	stmt.Exec()
	fmt.Println("Created table payments")
}