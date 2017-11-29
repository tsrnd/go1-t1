package controllers

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type (
	ProductController struct{}
)

func (hc ProductController) Product(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	context := Content{"Single Product"}

	// layout file must be the first parameter in ParseFiles!
	templates, err := template.ParseFiles(
		"views/layout/master.html",
		"views/layout/header.html",
		"views/layout/slider.html",
		"views/product/product.html",
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
