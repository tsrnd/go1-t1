package model

import (
	"log"
)

type User struct {
	Model
	Username string `schema:"username"`
	Fullname string `schema:"fullname"`
	Email    string `schema:"email"`
	Address  string `schema:"address"`
	Password string `schema:"password"`
	Orders   []Order
}

func GetAll() (users User, err error) {
	rows, err := DBCon.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return users, err
}

func GetUserByID(ID int64) (user User, err error) {
	err = DBCon.QueryRow("SELECT username, fullname, email, address, password FROM users where id = ?", ID).Scan(&user.Username, &user.Fullname, &user.Email, &user.Address, &user.Password)
	if err != nil {
		log.Fatal(err)
	}
	return user, err
}

func GetUserByUserName(username string) (user User, err error) {
	err = DBCon.QueryRow("SELECT username, fullname, email, address, password FROM users where id = ?", username).Scan(&user.Username, &user.Fullname, &user.Email, &user.Address, &user.Password)
	if err != nil {
		log.Fatal(err)
	}
	return user, err
}

func GetUserByEmail(email string) (user User, err error) {
	err = DBCon.QueryRow("SELECT email FROM users where email = ?", email).Scan(&user.Email)
	if err != nil {
		log.Fatal(err)
	}
	return user, err
}

func Create(username string, fullname string, email string, address string, password string) (user User, err error) {
	err = DBCon.QueryRow("INSERT INTO users(username,fullname,email,address,password) VALUES (?,?,?,?,?)", username, fullname, email, address, password).Scan(&user.ID)
	if err != nil {
		log.Fatal(err)
	}
	return user, err
}

func UpdateUser(ID int64, fullname string, address string, password string) (user User, err error) {
	err = DBCon.QueryRow("UPDATE user SET fullname =  ?, password=?, address = ? WHERE  id = ?", fullname, password, address, ID).Scan(&user.ID)
	if err != nil {
		log.Fatal(err)
	}
	return user, err
}
