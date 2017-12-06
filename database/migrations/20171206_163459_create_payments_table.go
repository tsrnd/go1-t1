package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreatePaymentsTable_20171206_163459 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreatePaymentsTable_20171206_163459{}
	m.Created = "20171206_163459"

	migration.Register("CreatePaymentsTable_20171206_163459", m)
}

// Run the migrations
func (m *CreatePaymentsTable_20171206_163459) Up() {
	m.SQL(`CREATE TABLE payments (
		id SERIAL PRIMARY KEY,	
		method VARCHAR(255) NOT NULL,
		status integer NOT NULL,
		address VARCHAR(255) NOT NULL,
		total integer NOT NULL,
		vat integer NOT NULL,
		order_id integer NOT NULL,
		created_at timestamp with time zone,
		updated_at timestamp with time zone,
		deleted_at timestamp with time zone
	);
	`)

}

// Reverse the migrations
func (m *CreatePaymentsTable_20171206_163459) Down() {
	m.SQL("DROP TABLE payments")

}
