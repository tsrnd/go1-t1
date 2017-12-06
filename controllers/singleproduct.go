package controllers

import (
	"goweb1/model"
	"html/template"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type (
	ProductController struct{}
)

func (hc ProductController) Product(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	ids, _ := strconv.ParseInt(id, 10, 64)

	product, _ := model.GetProductByID(ids)
	ProductRelated, _ := model.GetProductByCategory(product[0].CategoryId, ids)
	session, _ := store.Get(r, "session-id")
	username := session.Values["username"]
	cats, _ := model.GetAllCategory()
	data := map[string]interface{}{
		"Product":        product,
		"ProductRelated": ProductRelated,
		"cats":           cats,
		"name":           username,
	}

	// layout file must be the first parameter in ParseFiles!
	templates, err := template.ParseFiles(
		"views/layout/master.html",
		"views/layout/header.html",
		"views/product/product.html",
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
