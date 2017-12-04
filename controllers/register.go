package controllers

import (
	"html/template"
	"net/http"
	"goweb1/model"
	"github.com/julienschmidt/httprouter"
	"html"
	"fmt"
)

type (
	RegisterController struct{}
)

func (hc *RegisterController) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	sessionFash, _ := store.Get(r, "session-flash")	
	context := Data{ sessionFash.Flashes(),}
	sessionFash.Options.MaxAge = -1
	sessionFash.Save(r, w)
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
	if err := templates.Execute(w, context); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (hc RegisterController) RegisterPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	username := r.FormValue("username")
	fullname := html.EscapeString(r.FormValue("fullname"))
	mail := r.FormValue("mail")
	password, _ := HashPassword(r.FormValue("password"))
	address := html.EscapeString(r.FormValue("address"))
	v := new (Validator)
	if !v.ValidateUsername(username) || 
	   !v.ValidateUsername(username) || 
	   !v.ValidateEmail(mail) || 
	   !v.ValidatePassword(password) || 
	    v.ExistsEmail(mail) ||
	    v.ExistsUserName(username) {
		SessionFlash(v.err, w, r)
		http.Redirect(w, r, URL_REGISTER, http.StatusMovedPermanently)
		return
	}

	_, err := model.Create(username, fullname, mail, address, password)
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, URL_REGISTER, http.StatusMovedPermanently)
		return
	}
	http.Redirect(w, r, URL_LOGIN, http.StatusMovedPermanently)
}
