package migrations

import (
	_ "github.com/lib/pq"
	"fmt"
)

func PaymentsIndexOrderId() {
	stmt, err := DB.Prepare(`
		CREATE INDEX payments_order_id_index ON payments(order_id);
		`)
	checkErr(err)
	stmt.Exec()
	fmt.Println("Add index Key payments_order_id_index")
}