package model

import (
    "github.com/jinzhu/gorm"
)

type OrderItem struct {
    gorm.Model 
    Quantity  int `gorm:"not null"`
    Price  int `gorm:"not null"`
	ProductId  int `gorm:"not null"`
    OrderId	uint `gorm:"not null"`
    Product Product // OrderItem belong to Product
    Order Order // OrderItem belong to Order
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

func SaveOrderItem(quantity int, price int, product_id int, order_id uint) (*OrderItem, error ) {
	orderItem := OrderItem{ Quantity: quantity, Price:price , ProductId: product_id,OrderId: order_id}
    res := DB.Create(&orderItem)
    return &orderItem, res.Error
}