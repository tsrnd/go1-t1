package controllers

import (
	"goweb1/model"
	"html/template"
	"net/http"
)

type (
	HomeController struct{}
)

func (hc HomeController) Home(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-id")
	username := session.Values["username"]
	product, _ := model.GetAllProduct()
	cats, _ := model.GetAllCategory()
	hdata := map[string]interface{}{
		"product": product,
		"cats":    cats,
		"name":    username,
	}
	// layout file must be the first parameter in ParseFiles!
	templates, err := template.ParseFiles(
		"views/layout/master.html",
		"views/layout/header.html",
		"views/home/index.html",
		"views/layout/footer.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := templates.Execute(w, hdata); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
