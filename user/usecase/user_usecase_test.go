package usecase_test

import (
	models "goweb1/user"
	"goweb1/user/repository/mocks"
	userCase "goweb1/user/usecase"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FakeUserStruct struct {
	UserName string `faker:"username"`
	Email    string `faker:"email"`
}

func TestGetByID(t *testing.T) {
	userFake := FakeUserStruct{}
	err := faker.FakeData(&userFake)
	if err != nil {
		t.Error(err)
	}
	mockUserRepo := new(mocks.UserRepository)
	mockUser := models.User{
		ID:       1,
		Email:    userFake.Email,
		UserName: userFake.UserName,
	}
	mockUserRepo.On("GetByID", mock.AnythingOfType("int64")).Return(&mockUser, nil)
	defer mockUserRepo.AssertCalled(t, "GetByID", mock.AnythingOfType("int64"))
	u := userCase.NewUserUsecase(mockUserRepo)
	a, err := u.GetByID(mockUser.ID)
	assert.NoError(t, err)
	assert.NotNil(t, a)
	assert.Equal(t, mockUser, *a)
}
