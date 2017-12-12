package controllers

import (
	"goweb1/model"
	"html/template"
	"net/http"
)

type (
	NotFoundController struct{}
)

func (url_notfound NotFoundController) NotFound(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-id")
	username := session.Values["username"]
	cats, _ := model.GetAllCategory()
	data := map[string]interface{}{
		"cats": cats,
		"name": username,
	}

	// layout file must be the first parameter in ParseFiles!
	templates, err := template.ParseFiles(
		"views/layout/master.html",
		"views/layout/header.html",
		"views/page/notfound.html",
		"views/layout/footer.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := templates.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
