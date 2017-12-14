package controllers

import (
	"fmt"
	"goweb1/model"
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
	fmt.Println(product)
	ProductRelated, _ := model.GetProductByCategory(product.CategoryId, ids)
	session, _ := store.Get(r, "session-id")
	username := session.Values["username"]
	cats, _ := model.GetAllCategory()
	data := map[string]interface{}{
		"Product":        product,
		"ProductRelated": ProductRelated,
		"cats":           cats,
		"name":           username,
	}

	RenderTemplate(w, "views/product/product.html", data)
}
