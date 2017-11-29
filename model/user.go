package model

import (
    "github.com/jinzhu/gorm"
    // "fmt"
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
    var users User
    res := DB.Find(&users)
    return &users, res.Error
}

func GetOne(ID int) (*User, error ) {
	var user User
    res := DB.First(&user, ID)
    return &user, res.Error
}