package controllers

import (
	"goweb1/model"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type (
	NotFoundController struct{}
)

func (url_notfound NotFoundController) NotFound(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session, _ := store.Get(r, "session-id")
	username := session.Values["username"]
	cats, _ := model.GetAllCategory()
	data := map[string]interface{}{
		"cats": cats,
		"name": username,
	}
	RenderTemplate(w, "views/page/notfound.html", data)
}
