package controllers

import (
	"html/template"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"goweb1/model"
	"goweb1/messages"
)

type (
	LoginController struct{}
)
type Data struct {
	Errors []interface{}
}

// Login page
func (ctrl LoginController) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session, _ := store.Get(r, "session-id")
	sessionFash, _ := store.Get(r, "session-flash")
	context := Data{ sessionFash.Flashes()}
	sessionFash.Options.MaxAge = -1
	sessionFash.Save(r, w)
	if session.Values["username"] != nil {
		http.Redirect(w, r, URL_HOME, http.StatusMovedPermanently)
	}
	// layout file must be the first parameter in ParseFiles!
	templates, err := template.ParseFiles(
		"views/layout/master.html",
		"views/layout/header.html",
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
	user, _ := model.GetUserByUserName(username)
	v := new (Validator)

	if !v.ValidateUsername(username) {
		SessionFlash(v.err, w, r)
		http.Redirect(w, r, URL_LOGIN, http.StatusMovedPermanently)
		return
	}

	if user.Username == "" || !CheckPasswordHash(password, user.Password) {
		SessionFlash(messages.Error_username_or_password, w, r)
		http.Redirect(w, r, URL_LOGIN, http.StatusMovedPermanently)
		return
	}

	session.Values["username"] = user.Username
	session.Values["id"] = user.ID
	session.Save(r, w)
	http.Redirect(w, r, URL_HOME, http.StatusMovedPermanently)
}

func (hc *LoginController) LogOut(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session, _ := store.Get(r, "session-id")
	session.Values["username"] = nil
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, URL_HOME, http.StatusMovedPermanently)
}
