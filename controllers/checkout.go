package controllers

import (
	"fmt"
	"goweb1/model"
	"net/http"
	"strconv"

	"github.com/gorilla/csrf"
	"github.com/julienschmidt/httprouter"
)

const VAT = 5

type (
	CheckoutController struct{}
)

func (hc CheckoutController) Checkout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session, _ := store.Get(r, "session-id")
	username := session.Values["username"]
	id := session.Values["id"].(uint)
	order, _ := model.GetOrderByID(id)
	order_id := order.ID
	total := order.Total + VAT
	cats, _ := model.GetAllCategory()
	cdata := map[string]interface{}{
		"cats":           cats,
		"name":           username,
		"order_id":       order_id,
		"total":          total,
		csrf.TemplateTag: csrf.TemplateField(r),
	}
	RenderTemplate(w, "views/checkout/checkout.html", cdata)
}

func (hc CheckoutController) CheckoutPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	menthod := r.FormValue("menthod")
	address := r.FormValue("address")
	total, _ := strconv.ParseInt(r.FormValue("total"), 10, 64)
	vat, _ := strconv.ParseInt(r.FormValue("vat"), 10, 64)
	order_id, _ := strconv.ParseUint(r.FormValue("order_id"), 10, 64)
	_, err := model.CreatePayment(menthod, 1, address, total, vat, order_id)
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, URL_CHECKOUT, http.StatusMovedPermanently)
		return
	}
	http.Redirect(w, r, URL_HOME, http.StatusMovedPermanently)
}
