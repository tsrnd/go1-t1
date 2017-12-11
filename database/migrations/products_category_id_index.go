package migrations

import (
	_ "github.com/lib/pq"
	"fmt"
)

func ProductsIndexCategoryId() {
	stmt, err := DB.Prepare(`
		CREATE INDEX products_category_id_index ON products(category_id);
		`)
	checkErr(err)
	stmt.Exec()
	fmt.Println("Add index Key products_category_id_index")
}