package psql

import (
	"database/sql"
	"goweb1/services/crypto"
	model "goweb1/user"
	repo "goweb1/user/repository"
)

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(Conn *sql.DB) repo.UserRepository {

	return &userRepository{Conn}
}

func (m *userRepository) GetByID(id int64) (*model.User, error) {
	const query = `
    select
      id,
      email,
      username
    from
      users
    where
      id = $1
  `
	var user model.User
	err := m.DB.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.UserName)
	return &user, err
}

func (m *userRepository) GetByEmail(email string) (*model.User, error) {
	const query = `
    select
      id,
      email,
      username
    from
      users
    where
      email = $1
  `
	var user model.User
	err := m.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.UserName)
	return &user, err
}

// func (m *userRepository) GetPrivateUserDetailsByEmail(email string) (*model.PrivateUserDetails, error) {
// 	const query = `
//     select
//       id,
//       password,
//       salt
//     from
//       users
//     where
//       email = $1
//   `
// 	var u model.PrivateUserDetails
// 	err := m.DB.QueryRow(query, email).Scan(&u.ID, &u.Password, &u.Salt)
// 	return &u, err
// }

func (m *userRepository) Create(email, name, password string) (int64, error) {
	const query = `
    insert into users (
      email,
      name,
      password,
      salt
    ) values (
      $1,
      $2,
      $3,
      $4
    )
    returning id
  `
	salt := crypto.GenerateSalt()
	hashedPassword := crypto.HashPassword(password, salt)
	var id int64
	err := m.DB.QueryRow(query, email, name, hashedPassword, salt).Scan(&id)
	return id, err
}
