package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateOrderItemsTable_20171206_163348 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateOrderItemsTable_20171206_163348{}
	m.Created = "20171206_163348"

	migration.Register("CreateOrderItemsTable_20171206_163348", m)
}

// Run the migrations
func (m *CreateOrderItemsTable_20171206_163348) Up() {
	m.SQL(`CREATE TABLE order_items (
		id SERIAL PRIMARY KEY,
		quantity integer NOT NULL,
		price integer NOT NULL,
		product_id integer NOT NULL,
		order_id integer NOT NULL,
		created_at timestamp with time zone,
		updated_at timestamp with time zone,
		deleted_at timestamp with time zone
	);`)
}

// Reverse the migrations
func (m *CreateOrderItemsTable_20171206_163348) Down() {
	m.SQL("DROP TABLE order_items")

}
