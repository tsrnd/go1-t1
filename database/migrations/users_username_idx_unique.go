package migrations

import (
	_ "github.com/lib/pq"
	"fmt"
)

func UsersUniqueUsername() {
	stmt, err := DB.Prepare(`
		CREATE UNIQUE INDEX users_username_idx_unique on users (username);
		`)
	checkErr(err)
	stmt.Exec()
	fmt.Println("Add unique Key users_username_idx_unique")
}