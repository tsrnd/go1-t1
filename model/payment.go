package model

import (
    "github.com/jinzhu/gorm"
)

type Payment struct {
    gorm.Model
    ID    int   
    Method  string 
    Status  int 
	Address  string 
	Total	int
	Vat		int
	OrderId int
}
