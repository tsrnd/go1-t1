package model

import (
    "github.com/jinzhu/gorm"
)

type Product struct {
    gorm.Model
    ID    int   
    Name  string 
    Price  int 
    Image  string 
    Size  string 
	Color   string  
	CategoryId int 
}
