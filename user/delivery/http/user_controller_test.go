package http_test

import (
	"goweb1/user/usecase"
	"net/http"
	"testing"
)

func TestUserController_UserLogin(t *testing.T) {
	type fields struct {
		Usecase usecase.UserUsecase
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := &UserController{
				Usecase: tt.fields.Usecase,
			}
			ctrl.UserLogin(tt.args.w, tt.args.r)
		})
	}
}
