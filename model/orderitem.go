package model

import (
    "github.com/jinzhu/gorm"
)

type OrderItem struct {
    gorm.Model 
    Quantity  int `gorm:"not null"`
    Price  int `gorm:"not null"`
	ProductId  uint `gorm:"not null"`
	OrderId	uint `gorm:"not null"`
}

func GetAllOrderItems() (AllOrderItems []OrderItem , err error) {
    AllOrderItems = []OrderItem{}
	err = DB.Find(&AllOrderItems).Error
    return AllOrderItems, err
}

func GetOrderItemByOrder(order_id int64) (orderItems []OrderItem, err error) {
    orderItems = []OrderItem{}
    err = DB.Where("order_id = ?", order_id).Find(&orderItems).Error
    return orderItems, err
}