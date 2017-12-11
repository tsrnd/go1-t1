package migrations

import (
	_ "github.com/lib/pq"
	"fmt"
)

func CreateUsersTable() {
	stmt, err := DB.Prepare(`CREATE TABLE users(  
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) NOT NULL,
		fullname VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		address VARCHAR(255),
		password VARCHAR(255) NOT NULL,
		created_at timestamp with time zone,
		updated_at timestamp with time zone,
		deleted_at timestamp with time zone
	  );`)
	checkErr(err)
	stmt.Exec()
	fmt.Println("Created table users")
}