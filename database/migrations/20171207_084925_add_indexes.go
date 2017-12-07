package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddIndexes_20171207_084925 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddIndexes_20171207_084925{}
	m.Created = "20171207_084925"

	migration.Register("AddIndexes_20171207_084925", m)
}

// Run the migrations
func (m *AddIndexes_20171207_084925) Up() {
	m.SQL(`CREATE INDEX order_items_product_id_index ON order_items(product_id);`)
	m.SQL(`CREATE INDEX order_items_order_id_index ON order_items(order_id);`)
	m.SQL(`CREATE INDEX orders_user_id_index ON orders(user_id);`)
	m.SQL(`CREATE INDEX payments_order_id_index ON payments(order_id);`)
	m.SQL(`CREATE INDEX products_category_id_index ON products(category_id);`)
}

// Reverse the migrations
func (m *AddIndexes_20171207_084925) Down() {
	m.SQL(`DROP INDEX order_items_product_id_index;`)
	m.SQL(`DROP INDEX order_items_order_id_index;`)
	m.SQL(`DROP INDEX orders_user_id_index;`)
	m.SQL(`DROP INDEX payments_order_id_index;`)
	m.SQL(`DROP INDEX products_category_id_index;`)

}
