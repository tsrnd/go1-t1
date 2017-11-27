package database

import "github.com/jinzhu/gorm"

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/goweb1?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	return db
}
