package controllers

import (
	"fmt"
	"goweb1/model"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type (
	RegisterController struct{}
)

func (hc *RegisterController) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session, _ := store.Get(r, "session-id")
	username := session.Values["username"]
	fmt.Println(username)
	cats, _ := model.GetAllCategory()
	rdata := map[string]interface{}{
		"cats": cats,
		"name": username,
	}

	// layout file must be the first parameter in ParseFiles!
	templates, err := template.ParseFiles(
		"views/layout/master.html",
		"views/layout/header.html",
		"views/register/register.html",
		"views/layout/footer.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := templates.Execute(w, rdata); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (hc RegisterController) RegisterPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	username := r.FormValue("username")
	fullname := r.FormValue("fullname")
	mail := r.FormValue("mail")
	password, _ := HashPassword(r.FormValue("password"))
	address := r.FormValue("address")
	model.Create(username, fullname, mail, address, password)
	http.Redirect(w, r, URL_LOGIN, http.StatusMovedPermanently)
}
