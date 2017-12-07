package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateCategoriesTable_20171206_163314 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateCategoriesTable_20171206_163314{}
	m.Created = "20171206_163314"

	migration.Register("CreateCategoriesTable_20171206_163314", m)
}

// Run the migrations
func (m *CreateCategoriesTable_20171206_163314) Up() {
	m.SQL(`CREATE TABLE categories (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		created_at timestamp with time zone,
		updated_at timestamp with time zone,
		deleted_at timestamp with time zone
	);`)

}

// Reverse the migrations
func (m *CreateCategoriesTable_20171206_163314) Down() {
	m.SQL("DROP TABLE categories")

}
