package main

import (
	"github.com/astaxie/beego/migration"
	_ "github.com/lib/pq"
)

// DO NOT MODIFY
type CreateUsersTable_20171206_140924 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUsersTable_20171206_140924{}
	m.Created = "20171206_140924"

	migration.Register("CreateUsersTable_20171206_140924", m)
}

// Run the migrations
func (m *CreateUsersTable_20171206_140924) Up() {
	m.SQL(`CREATE TABLE users(  
			id SERIAL PRIMARY KEY,
			username VARCHAR NOT NULL,
			fullname VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			address VARCHAR(255),
			password VARCHAR(255) NOT NULL,
			created_at timestamp NOT NULL,
			updated_at timestamp NOT NULL,
			deleted_at timestamp NOT NULL
		  );`)
}

// Reverse the migrations
func (m *CreateUsersTable_20171206_140924) Down() {
	m.SQL("DROP TABLE users")

}