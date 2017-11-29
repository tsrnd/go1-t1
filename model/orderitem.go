package model

import (
    "github.com/jinzhu/gorm"
)

type OrderItem struct {
    gorm.Model
    ID    int   
    Quantity  int 
    Price  int 
	ProductId  int 
	OrderId	int
}
