package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

const (
	URL_HOME     = "/"
	URL_LOGIN    = "/login"
	URL_NOTFOUND = "/notfound"
	URL_REGISTER = "/register"
	URL_CHECKOUT = "/checkout"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func SessionFlash(err string, w http.ResponseWriter, r *http.Request) {
	sessionFash, _ := store.Get(r, "session-flash")
	sessionFash.AddFlash(err)
	sessionFash.Save(r, w)
}

func CoverInterfaceToString(inter []interface{}) []string {

	s := make([]string, len(inter))
	for i, v := range inter {
		s[i] = fmt.Sprint(v)
	}
	return s
}

func AuthIsLogin(r *http.Request) bool {
	session, _ := store.Get(r, "session-id")
	return session.Values["username"] != nil
}

func RenderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	templates, err := template.ParseFiles(
		"views/layout/master.html",
		"views/layout/header.html",
		templateFile,
		"views/layout/footer.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := templates.Execute(w, templateData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
