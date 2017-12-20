package infrastructure

import (
	"github.com/jinzhu/gorm"
	// blank import.
	_ "github.com/lib/pq"
)

// SQL struct.
type SQL struct {
	Database *gorm.DB
}

// NewSQL returns new SQL.
func NewSQL() *SQL {
	dbms := GetConfigString("database.dbms")
	host := GetConfigString("database.host")
	user := GetConfigString("database.user")
	pass := GetConfigString("database.pass")
	name := GetConfigString("database.name")

	connect := "host=" + host + " user=" + user + " dbname=" + name + " sslmode=disable password=" + pass
	db, err := gorm.Open(dbms, connect)
	// Disable table name's pluralization globally
	// if set this to true, `User`'s default table name will be `user`, table name setted with `TableName` won't be affected
	db.SingularTable(true)
	if err != nil {
		panic(err)
	}

	return &SQL{db}
}
