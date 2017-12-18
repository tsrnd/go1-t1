package repository_test

import (
	"fmt"
	"testing"

	userRepo "goweb1/user/repository/psql"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/bxcodec/faker"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

type FakeUserStruct struct {
	UserName string `faker:"username"`
	Email    string `faker:"email"`
	Password string `faker:"password"`
}

func TestGetByID(t *testing.T) {
	userFake := FakeUserStruct{}
	err := faker.FakeData(&userFake)
	if err != nil {
		fmt.Println(err)
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{
		"id", "email", "username"}).
		AddRow(1, userFake.Email, userFake.UserName)
	query := "select id, email, username from users where id = ?"
	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)
	repo := userRepo.NewUserRepository(db)
	num := int64(1)
	user, err := repo.GetByID(num)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}
