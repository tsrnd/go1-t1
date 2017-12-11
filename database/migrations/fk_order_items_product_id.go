package migrations

import (
	_ "github.com/lib/pq"
	"fmt"
)

func OrderItemsForeignKeysProductId() {
	stmt, err := DB.Prepare(`
		ALTER TABLE order_items 
		ADD CONSTRAINT fk_order_items_product_id
		FOREIGN KEY (product_id) 
		REFERENCES products(id);
		`)
	checkErr(err)
	stmt.Exec()
	fmt.Println("Add Foreign Keys fk_order_items_product_id")
}