package model

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	Total      int64 `gorm:"not null"`
	Status     int   `gorm:"not null"`
	UserId     uint
	OrderItems []OrderItem // Order hasmany OrderItem
	Payments   []Payment
	User       User // Order belong to User
}

func GetOrderByID(ID uint) (*Order, error) {
	var order Order
	erro := DB.Select("id,total,status").Where("user_id = ?", ID).Last(&order)
	return &order, erro.Error
}

func SaveOrder(total int64, user_id uint) (*Order, error) {
	order := Order{Total: total, UserId: user_id, Status: 0}
	res := DB.Create(&order)
	return &order, res.Error
}

func GetIdOrder() (*Order, error) {
	var orders Order
	res := DB.Select("id").Last(&orders)
	return &orders, res.Error
}
