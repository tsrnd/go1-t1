package model

import (
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    ID    int   
    Username  string 
    Fullname  string 
    Mail  string 
    Address  string 
    Password   string   
}

func GetAll() (*User, error) {
    var user User
    res := DB.Find(&user)
    return &user, res.Error
}