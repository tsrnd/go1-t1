package model

import (
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Age   int    `json:"age"`
}

func GetAll() (*User, error) {
    var user User
    res := DB.Find(&user)
    return &user, res.Error
}