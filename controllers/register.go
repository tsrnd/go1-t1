package controllers

import (
	"html/template"
	"net/http"
	"fmt"
	"goweb1/model"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"

)

type (
	RegisterController struct{}
)

func (hc *RegisterController) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	context := Content{"Order"}

	// layout file must be the first parameter in ParseFiles!
	templates, err := template.ParseFiles(
		"views/layout/master.html",
		"views/layout/header.html",
		"views/layout/slider.html",
		"views/register/register.html",
		"views/layout/footer.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := templates.Execute(w, context); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (hc RegisterController) RegisterPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	fmt.Println(r.FormValue("username"))
	username := r.FormValue("username")
	fullname := r.FormValue("fullname")
	mail :=r.FormValue("mail")
	password, _ := HashPassword(r.FormValue("password"))
	address :=r.FormValue("address")
	model.Create(username, fullname, mail, address, password)
	http.Redirect(w, r, "/", 301)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}