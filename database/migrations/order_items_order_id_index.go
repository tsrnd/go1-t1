package migrations

import (
	_ "github.com/lib/pq"
	"fmt"
)

func OrderItemsIndexOrderId() {
	stmt, err := DB.Prepare(`
		CREATE INDEX order_items_order_id_index ON order_items(order_id);
		`)
	checkErr(err)
	stmt.Exec()
	fmt.Println("Add index Key order_items_order_id_index")
}