package repository

import model "goweb1/user"

// UserRepository interface
type UserRepository interface {
	GetByID(id int64) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	//GetPrivateDetailsByEmail(email string) (*model.PrivateUserDetails, error)
	Create(email, name, password string) (int64, error)
}
