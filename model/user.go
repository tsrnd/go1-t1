package model

import (

    "github.com/jinzhu/gorm"
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

func GetUserByUserName(username string)  (*User, error ){
    var user User
    res := DB.Select("username, password").Where("username = ?", username).First(&user)
    return &user, res.Error
}

func Create(username string, fullname string, email string, address string, password string) (*User, error ) {
	user := User{Username: username, Fullname: fullname, Email: email, Address: address, Password: password}
    res := DB.Create(&user)
    return &user, res.Error
}

