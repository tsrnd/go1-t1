package model

import "fmt"

type User struct {
	Id       int
	Username string `gorm:"not null"`
	Fullname string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Address  string `gorm:"not null"`
	Password string `gorm:"not null"`
	Orders   []Order
}

// func GetAll() (*User, error) {
//     var users User
//     res := DB.Find(&users)
//     return &users, res.Error
// }

// func GetUserByID(ID int64) (user []User, err error) {
//     user = []User {}
//     err = DB.First(&user, ID).Error
//     return user, err
// }

func GetUserByUserName(username string) (user User, err error) {
	user = User{}
	err = DB.QueryRow("select id,username,password from users where username = $1 order by id desc limit 1", username).Scan(&user.Id, &user.Username, &user.Password)
	return
}

// func GetUserByEmail(email string)  (*User, error ){
//     var user User
//     res := DB.Select("email").Where("email = ?", email).First(&user)
//     return &user, res.Error
// }

func Create(username string, fullname string, email string, address string, password string) (err error) {
	statement := "insert into users (username, fullname, email, address, password) values ($1, $2, $3, $4, $5) returning id"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	id := 0
	err = stmt.QueryRow(username, fullname, email, address, password).Scan(&id)
	fmt.Println(err)
	return

}

// func UpdateUser(ID int64, fullname string,  address string, password string) (*User, error) {
//     var user User
//     DB.First(&user, ID)
//     user.Fullname = fullname
//     user.Password = password
//     user.Address = address
//     res := DB.Save(&user)
//     return &user, res.Error
// }
