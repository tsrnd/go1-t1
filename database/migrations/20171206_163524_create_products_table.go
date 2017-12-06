package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateProductsTable_20171206_163524 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateProductsTable_20171206_163524{}
	m.Created = "20171206_163524"

	migration.Register("CreateProductsTable_20171206_163524", m)
}

// Run the migrations
func (m *CreateProductsTable_20171206_163524) Up() {
	m.SQL(`CREATE TABLE products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		price integer NOT NULL,
		image VARCHAR(255) NOT NULL,
		size integer NOT NULL,
		color VARCHAR(255) NOT NULL,
		category_id integer NOT NULL,
		created_at timestamp with time zone,
		updated_at timestamp with time zone,
		deleted_at timestamp with time zone
	);
	`)

}

// Reverse the migrations
func (m *CreateProductsTable_20171206_163524) Down() {
	m.SQL("DROP TABLE products")

}
