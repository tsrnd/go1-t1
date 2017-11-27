package model

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func SetDatabase(database *gorm.DB) {
	DB = database
}
