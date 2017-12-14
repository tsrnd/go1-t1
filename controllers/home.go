package controllers

import (
	"goweb1/model"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type (
	HomeController struct{}
)

func (hc HomeController) Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
	RenderTemplate(w, "views/home/index.html", hdata)
}
