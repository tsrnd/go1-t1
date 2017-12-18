package http_test

import (
	"goweb1/config"
	deli "goweb1/user/delivery/http"
	repos "goweb1/user/repository/psql"
	"goweb1/user/usecase"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserController_UserLogin(t *testing.T) {
	userRepo := repos.NewFakeUserRepository(config.DB())
	userUsecase := usecase.NewUserUsecase(userRepo)
	type fields struct {
		Usecase usecase.UserUsecase
	}
	type args struct {
		rr  *httptest.ResponseRecorder
		req *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			"test login success",
			fields{Usecase: userUsecase},
			args{
				rr: httptest.NewRecorder(),
				req: func() (req *http.Request) {
					req, err := http.NewRequest("GET", "/users", nil)
					if err != nil {
						t.Fatal(err)
					}
					return req
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := &deli.UserController{
				Usecase: tt.fields.Usecase,
			}
			handler := http.HandlerFunc(ctrl.UserLogin)
			handler.ServeHTTP(tt.args.rr, tt.args.req)
			if tt.args.rr.Code != 200 {
				t.Errorf("loi")
			}
		})
	}
}
