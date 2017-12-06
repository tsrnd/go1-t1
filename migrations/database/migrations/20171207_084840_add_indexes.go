package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddIndexes_20171207_084840 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddIndexes_20171207_084840{}
	m.Created = "20171207_084840"

	migration.Register("AddIndexes_20171207_084840", m)
}

// Run the migrations
func (m *AddIndexes_20171207_084840) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *AddIndexes_20171207_084840) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
