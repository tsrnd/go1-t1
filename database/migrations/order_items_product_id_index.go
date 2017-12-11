package migrations

import (
	_ "github.com/lib/pq"
	"fmt"
)

func OrderItemsIndexProductId() {
	stmt, err := DB.Prepare(`
		CREATE INDEX order_items_product_id_index ON order_items(product_id);
		`)
	checkErr(err)
	stmt.Exec()
	fmt.Println("Add index Key order_items_product_id_index")
}