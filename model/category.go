package model

import (
    "github.com/jinzhu/gorm"
)

type Category struct {
    gorm.Model
    Name  string `gorm:"not null"`
    Products   []Product // Category has many product
}

func GetAllCategory() (AllCategory []Category, erro error) {
    AllCategory = []Category{}
	erro = DB.Find(&AllCategory).Error
    return AllCategory, erro
}