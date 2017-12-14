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
	user = User{}
	err = DBCon.QueryRow("select id,username,password from users where username = $1 order by id desc limit 1", username).Scan(&user.ID, &user.Username, &user.Password)
	return
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	err = DBCon.QueryRow("select email from users where email = $1 order by id desc limit 1", email).Scan(&user.Email)
	return
}

func CreateUser(username string, fullname string, email string, address string, password string) (err error) {
	statement := "insert into users (username, fullname, email, address, password) values ($1, $2, $3, $4, $5) returning id"
	stmt, err := DBCon.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	id := 0
	err = stmt.QueryRow(username, fullname, email, address, password).Scan(&id)
	return

}

func UpdateUser(ID int64, fullname string, address string, password string) (user User, err error) {
	err = DBCon.QueryRow("UPDATE user SET fullname =  ?, password=?, address = ? WHERE  id = ?", fullname, password, address, ID).Scan(&user.ID)
	if err != nil {
		log.Fatal(err)
	}
	return user, err
}
