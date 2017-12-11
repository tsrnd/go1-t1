package migrations

import (
	_ "github.com/lib/pq"
	"fmt"
)

func OrdersIndexUserId() {
	stmt, err := DB.Prepare(`
		CREATE INDEX orders_user_id_index ON orders(user_id);
		`)
	checkErr(err)
	stmt.Exec()
	fmt.Println("Add index Key orders_user_id_index")
}