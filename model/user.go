package model

import (
    "github.com/jinzhu/gorm"
    // "fmt"
)

type User struct {
    gorm.Model 
    Username  string `gorm:"not null"`
    Fullname  string `gorm:"not null"`
    Email  string `gorm:"not null"`
    Address  string `gorm:"not null"`
    Password   string `gorm:"not null"`
}
func GetAll() (*User, error) {
    var users User
    res := DB.Find(&users)
    return &users, res.Error
}

func GetOne(ID int) (*User, error ) {
	var user User
    res := DB.First(&user, ID)
    return &user, res.Error
}