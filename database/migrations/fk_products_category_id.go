package migrations

import (
	_ "github.com/lib/pq"
	"fmt"
)

func ProductsForeignKeysCategoryId() {
	stmt, err := DB.Prepare(`
		ALTER TABLE products 
		ADD CONSTRAINT fk_products_category_id
		FOREIGN KEY (category_id) 
		REFERENCES categories(id);
		`)
	checkErr(err)
	stmt.Exec()
	fmt.Println("Add Foreign Keys fk_products_category_id")
}