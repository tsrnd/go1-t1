package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddForeign_20171206_171257 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddForeign_20171206_171257{}
	m.Created = "20171206_171257"

	migration.Register("AddForeign_20171206_171257", m)
}

// Run the migrations
func (m *AddForeign_20171206_171257) Up() {
	m.SQL(`ALTER TABLE order_items 
		ADD CONSTRAINT fk_order_items_product_id
		FOREIGN KEY (product_id) 
		REFERENCES products(id);`)
	m.SQL(`ALTER TABLE order_items 
		ADD CONSTRAINT fk_order_items_order_id
		FOREIGN KEY (order_id) 
		REFERENCES orders(id);`)
	m.SQL(`ALTER TABLE orders 
		ADD CONSTRAINT fk_orders_user_id
		FOREIGN KEY (user_id) 
		REFERENCES users(id);`)
	m.SQL(`ALTER TABLE payments 
		ADD CONSTRAINT fk_payments_order_id
		FOREIGN KEY (order_id) 
		REFERENCES orders(id);`)
	m.SQL(`ALTER TABLE products 
		ADD CONSTRAINT fk_products_category_id
		FOREIGN KEY (category_id) 
		REFERENCES categories(id);`)
}

// Reverse the migrations
func (m *AddForeign_20171206_171257) Down() {
	m.SQL(`ALTER TABLE order_items DROP CONSTRAINT fk_order_items_product_id;`)
	m.SQL(`ALTER TABLE order_items DROP CONSTRAINT fk_order_items_order_id;`)
	m.SQL(`ALTER TABLE orders DROP CONSTRAINT fk_orders_user_id;`)
	m.SQL(`ALTER TABLE payments DROP CONSTRAINT fk_payments_order_id;`)
	m.SQL(`ALTER TABLE products DROP CONSTRAINT fk_products_category_id;`)
}
