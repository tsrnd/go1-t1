package controllers

import (
	"html/template"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"goweb1/model"
)

type (
	LoginController struct{}
)

// Login page
func (ctrl LoginController) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session, _ := store.Get(r, "session-id")
	context := Content{"Order"}

	if session.Values["username"] != nil {
		http.Redirect(w, r, URL_HOME, http.StatusMovedPermanently)
	}
	// layout file must be the first parameter in ParseFiles!
	templates, err := template.ParseFiles(
		"views/layout/master.html",
		"views/layout/header.html",
		"views/layout/slider.html",
		"views/login/login.html",
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
// ProcessLogin process login and session
func (ctrl LoginController) ProcessLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session, _ := store.Get(r, "session-id")
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	user, _ := model.CheckLogin(username, password)
	if user.Username== "" {
		http.Redirect(w, r, URL_LOGIN, http.StatusMovedPermanently)
	} 
	session.Values["username"] = user.Username
	session.Save(r, w)
	http.Redirect(w, r, URL_HOME, http.StatusMovedPermanently)
}
