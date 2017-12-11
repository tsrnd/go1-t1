package migrations

import (
	_ "github.com/lib/pq"
	"fmt"
)

func OrdersForeignKeysUserId() {
	stmt, err := DB.Prepare(`
		ALTER TABLE orders 
		ADD CONSTRAINT fk_orders_user_id
		FOREIGN KEY (user_id) 
		REFERENCES users(id);
		`)
	checkErr(err)
	stmt.Exec()
	fmt.Println("Add Foreign Keys fk_orders_user_id")
}