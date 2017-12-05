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
    Orders []Order
}


func GetAll() (*User, error) {
    var users User
    res := DB.Find(&users)
    return &users, res.Error
}

func GetUserByID(ID int64) (user []User, err error) {
    user = []User {}
    err = DB.First(&user, ID).Error
    return user, err
}

func GetUserByUserName(username string)  (*User, error ){
    var user User
    res := DB.Select("id,username, password").Where("username = ?", username).First(&user)
    return &user, res.Error
}

func GetUserByEmail(email string)  (*User, error ){
    var user User
    res := DB.Select("email").Where("email = ?", email).First(&user)
    return &user, res.Error
}

func Create(username string, fullname string, email string, address string, password string) (*User, error ) {
	user := User{Username: username, Fullname: fullname, Email: email, Address: address, Password: password}
    res := DB.Create(&user)
    return &user, res.Error
}

func UpdateUser(ID int64, fullname string,  address string, password string) (*User, error) {
    var user User
    DB.First(&user, ID)
    user.Fullname = fullname
    user.Password = password
    user.Address = address
    res := DB.Save(&user)
    return &user, res.Error
}