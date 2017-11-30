package controllers

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type (
	RegisterController struct{}
)

func (hc RegisterController) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {


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

	if err := templates.Execute(w, ""); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
