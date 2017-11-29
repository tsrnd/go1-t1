package model

import (
    "github.com/jinzhu/gorm"
)

type Product struct {
    gorm.Model
    Name  string `gorm:"not null"`
    Price  int `gorm:"not null"`
    Image  string `gorm:"not null"`
    Size  string `gorm:"not null"`
	Color   string `gorm:"not null"`
	CategoryId uint `gorm:"not null"`
}
