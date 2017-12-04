package model

import (
    "github.com/jinzhu/gorm"
)

type Category struct {
    gorm.Model
    Name  string `gorm:"not null"`
}

func GetAllCategory() (AllCategory []Category, erro error) {
    AllCategory = []Category{}
	erro = DB.Find(&AllCategory).Error
    return AllCategory, erro
}