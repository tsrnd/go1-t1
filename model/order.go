package model

import (
    "github.com/jinzhu/gorm"
)

type Order struct {
    gorm.Model
    Total  int64 `gorm:"not null"`
    Status  int `gorm:"not null"`
    UserId  uint 
}

func SaveOrder(total int64, user_id uint) (*Order, error ) {
	order := Order{Total: total, UserId: user_id, Status: 0}
    res := DB.Create(&order)
    return &order, res.Error
}
