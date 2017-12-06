package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateOrdersTable_20171206_163418 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateOrdersTable_20171206_163418{}
	m.Created = "20171206_163418"

	migration.Register("CreateOrdersTable_20171206_163418", m)
}

// Run the migrations
func (m *CreateOrdersTable_20171206_163418) Up() {
	m.SQL(`CREATE TABLE orders(
		id SERIAL PRIMARY KEY,
		total integer NOT NULL,
		status boolean NOT NULL,
		user_id integer,
		created_at timestamp with time zone,
		updated_at timestamp with time zone,
		deleted_at timestamp with time zone
	);`)

}

// Reverse the migrations
func (m *CreateOrdersTable_20171206_163418) Down() {
	m.SQL("DROP TABLE orders")

}
