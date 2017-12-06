package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddUniques_20171207_090221 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddUniques_20171207_090221{}
	m.Created = "20171207_090221"

	migration.Register("AddUniques_20171207_090221", m)
}

// Run the migrations
func (m *AddUniques_20171207_090221) Up() {
	m.SQL(`CREATE UNIQUE INDEX users_username_idx_unique on users (username);`) 
	m.SQL(`CREATE UNIQUE INDEX users_email_idx_unique on users (email);`) 

}

// Reverse the migrations
func (m *AddUniques_20171207_090221) Down() {
	m.SQL("DROP INDEX users_username_idx_unique;")
	m.SQL("DROP INDEX users_email_idx_unique;")
}
