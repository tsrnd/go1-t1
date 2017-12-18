package http

import (
	"fmt"
	"net/http"

	"goweb1/user/usecase"

	"github.com/go-chi/chi"
)

// UserController type
type UserController struct {
	Usecase usecase.UserUsecase
}

// NewUserController func
func NewUserController(r chi.Router, uc usecase.UserUsecase) *UserController {
	handler := &UserController{
		Usecase: uc,
	}
	// r.POST("/users", handler.UserRegister)
	// r.POST("/auth", handler.UserLogin)
	r.Get("/users", handler.UserLogin)
	return handler
}

func (ctrl *UserController) UserLogin(w http.ResponseWriter, r *http.Request) {
	id := int64(1)

	user, _ := ctrl.Usecase.GetByID(id)
	fmt.Println(user)
	// // p := map[string]string{
	// // 	"token": "test",
	// // }
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(user)
}

// UserRegister func
// func (ctrl *UserController) UserRegister(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.Error(w, "Not found", http.StatusNotFound)
// 		return
// 	}
// 	decoder := json.NewDecoder(r.Body)
// 	var rr requests.UserRegisterRequest
// 	err := decoder.Decode(&rr)
// 	if err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}
// 	id, err := repositories.CreateUser(ctrl.DB, rr.Email, rr.Name, rr.Password)
// 	if err != nil {
// 		log.Fatalf("Add user to database error: %s", err)
// 		http.Error(w, "", http.StatusInternalServerError)
// 		return
// 	}
// 	token, err := crypto.GenerateToken()
// 	if err != nil {
// 		log.Fatalf("Generate token Error: %s", err)
// 		http.Error(w, "", http.StatusInternalServerError)
// 		return
// 	}
// 	oneMonth := time.Duration(60*60*24*30) * time.Second
// 	err = ctrl.Cache.Set(fmt.Sprintf("token_%s", token), strconv.Itoa(id), oneMonth)
// 	if err != nil {
// 		log.Fatalf("Add token to redis Error: %s", err)
// 		http.Error(w, "", http.StatusInternalServerError)
// 		return
// 	}
// 	p := map[string]string{
// 		"token": token,
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(p)
// }
