package model

import (
    "github.com/jinzhu/gorm"
)

type Order struct {
    gorm.Model
    Total  int `gorm:"not null"`
    Status  int `gorm:"not null"`
    UserId  uint 
}
