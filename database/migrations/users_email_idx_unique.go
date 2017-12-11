package migrations

import (
	_ "github.com/lib/pq"
	"fmt"
)

func UsersUniqueEmail() {
	stmt, err := DB.Prepare(`
		CREATE UNIQUE INDEX users_email_idx_unique on users (email);
		`)
	checkErr(err)
	stmt.Exec()
	fmt.Println("Add unique Key users_email_idx_unique")
}