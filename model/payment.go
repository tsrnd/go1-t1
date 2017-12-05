package model

import (
    "github.com/jinzhu/gorm"
)

type Payment struct {
    gorm.Model  
    Method  string `gorm:"not null"`
    Status  int `gorm:"not null"`
	Address  string `gorm:"not null"`
	Total	int `gorm:"not null"`
	Vat		int `gorm:"not null"`
	OrderId uint `gorm:"not null"`
	Orders []Order 
}
