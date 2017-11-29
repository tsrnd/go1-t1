package controllers

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type (
	LoginController struct{}
)

func (hc LoginController) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	context := Content{"Order"}

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
