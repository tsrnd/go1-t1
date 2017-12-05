package model

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	Total      int `gorm:"not null"`
	Status     int `gorm:"not null"`
	UserId     uint
	OrderItems []OrderItem // Order hasmany OrderItem
	Payments   []Payment
	User       User // Order belong to User
}

func GetOrderByID(ID uint) (*Order, error) {
	var order Order
	erro := DB.Select("id,total,status").Where("user_id = ?", ID).First(&order)
	return &order, erro.Error
}
