package migrations

import (
	_ "github.com/lib/pq"
	"fmt"
)

func PaymentsForeignKeysOrderId() {
	stmt, err := DB.Prepare(`
		ALTER TABLE payments 
		ADD CONSTRAINT fk_payments_order_id
		FOREIGN KEY (order_id) 
		REFERENCES orders(id);
		`)
	checkErr(err)
	stmt.Exec()
	fmt.Println("Add Foreign Keys fk_payments_order_id")
}