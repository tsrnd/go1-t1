package database


import (
	"github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
	"os"
	"fmt"
)

func ConnectDB() *gorm.DB {
	host := os.Getenv("DB_CONNECTION")
	address := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_CHARSET"),
		os.Getenv("DB_PARSETIME"),
	)
	db, err := gorm.Open(host, address)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	return db
}
