package model

import (
	"github.com/jinzhu/gorm"
)

type Payment struct {
	gorm.Model
	Method  string `gorm:"not null"`
	Status  int    `gorm:"not null"`
	Address string `gorm:"not null"`
	Total   int64  `gorm:"not null"`
	Vat     int64  `gorm:"not null"`
	OrderId uint64 `gorm:"not null"`
	Orders  []Order
}

// func CreatePayment(method string, status int, address string, total int64, vat int64, order_id uint64) (*Payment, error) {
// 	payment := Payment{Method: method, Status: status, Address: address, Total: total, Vat: vat, OrderId: order_id}
// 	res := DB.Create(&payment)
// 	return &payment, res.Error
// }
