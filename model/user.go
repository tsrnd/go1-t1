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

func GetAll() (users []User, err error) {
	rows, err := DBCon.Query("SELECT username, fullname, email,address,password FROM users")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.ID, &user.Username, &user.Fullname, &user.Email, &user.Address, &user.Password)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	defer rows.Close()
	return users, err
}

func GetUserByID(ID int64) (user User, err error) {
	err = DBCon.QueryRow("SELECT username, fullname, email, address, password FROM users where id = $1", ID).Scan(&user.Username, &user.Fullname, &user.Email, &user.Address, &user.Password)
	if err != nil {
		log.Fatal(err)
	}
	return user, err
}

func GetUserByUserName(username string) (user User, err error) {
	err = DBCon.QueryRow("SELECT username, fullname, email, address, password FROM users where id = $1", username).Scan(&user.Username, &user.Fullname, &user.Email, &user.Address, &user.Password)
	if err != nil {
		log.Fatal(err)
	}
	return user, err
}

func GetUserByEmail(email string) (user User, err error) {
	err = DBCon.QueryRow("SELECT email FROM users where email = $1", email).Scan(&user.Email)
	if err != nil {
		log.Fatal(err)
	}
	return user, err
}

func (user *User) Create() (err error) {
	statement := "insert into users (username,fullname,email,address,password) VALUES ($1,$2,$3,$4,$5) returning id"
	stmt, err := DBCon.Prepare(statement)
	if err != nil {
		log.Fatal(err)
	}
	err = stmt.QueryRow(user.Username, user.Fullname, user.Email, user.Address, user.Password).Scan(&user.ID)
	return
}

func UpdateUser(ID int64, fullname string, address string, password string) (user User, err error) {
	err = DBCon.QueryRow("UPDATE user SET fullname = $1, password=$2, address = $3 WHERE  id = $4", fullname, password, address, ID).Scan(&user.ID)
	if err != nil {
		log.Fatal(err)
	}
	return user, err
}
