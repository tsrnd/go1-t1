package migrations

import (
	_ "github.com/lib/pq"
	"fmt"
)

func OrderItemsForeignKeysOrderId() {
	stmt, err := DB.Prepare(`
		ALTER TABLE order_items 
		ADD CONSTRAINT fk_order_items_order_id
		FOREIGN KEY (order_id) 
		REFERENCES orders(id);
		`)
	checkErr(err)
	stmt.Exec()
	fmt.Println("Add Foreign Keys fk_order_items_order_id")
}