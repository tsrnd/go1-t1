package model

import (
    "github.com/jinzhu/gorm"
)

type Order struct {
    gorm.Model
    Total  int `gorm:"not null"`
    Status  int `gorm:"not null"`
    UserId  uint 
    OrderItems []OrderItem // Order hasmany OrderItem
    Payments []Payment 
    User User // Order belong to User
}
