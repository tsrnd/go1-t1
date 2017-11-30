package model

import (
    "github.com/jinzhu/gorm"
)

type Product struct {
    gorm.Model
    Name  string `gorm:"not null"`
    Price  int `gorm:"not null"`
    Image  string `gorm:"not null"`
    Size  string `gorm:"not null"`
	Color   string `gorm:"not null"`
	CategoryId uint `gorm:"not null"`
}
func GetAllProduct() (AllProducts []Product, erro error) {
    AllProducts = []Product{}
	erro = DB.Find(&AllProducts).Error
    return AllProducts, erro
}

func GetProductByCategory(category_id uint, ID int64) (products []Product, erro error) {
    products = []Product{}
    erro = DB.Where("category_id = ?&& id != ?", category_id,ID).Find(&products).Error
    return products, erro
} 

func GetProductByID(ID int64) (product []Product, erro error) {
    product = []Product {}
    erro = DB.First(&product, ID).Error
    return product, erro
}


