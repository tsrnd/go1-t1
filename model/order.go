package model

import (
    "github.com/jinzhu/gorm"
)

type Order struct {
    gorm.Model
    ID    int   
    Total  int 
    Status  int 
    UserId  int 
}
