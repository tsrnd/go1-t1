package database


import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
	"fmt"
)

func ConnectDB() *gorm.DB {
	dbConnect := os.Getenv("DB_CONNECTION")
	address := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"),
		os.Getenv("SSLMODE"),
		os.Getenv("DB_PASSWORD"),
	)
	db, err := gorm.Open(dbConnect, address)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	return db
}
