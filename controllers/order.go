package controllers

import (
	"goweb1/model"
	"net/http"
	"strconv"

	"github.com/gorilla/csrf"
	"github.com/julienschmidt/httprouter"
)

type (
	OrderController struct{}
)

func (order_controller OrderController) Order(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	session, _ := store.Get(r, "session-id")
	username := session.Values["username"]
	cats, _ := model.GetAllCategory()
	data := map[string]interface{}{
		"cats": cats,
		"name": username,

		csrf.TemplateTag: csrf.TemplateField(r),
	}
	RenderTemplate(w, "views/order/order.html", data)
}

func (order_controller OrderController) SaveOrder(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// if err := r.ParseForm(); err != nil {
	// }
	session, _ := store.Get(r, "session-id")
	id_user := session.Values["id"].(uint)
	total := r.FormValue("totals")
	totals, _ := strconv.ParseInt(total, 10, 64)

	model.SaveOrder(totals, 0, id_user)
	orders, _ := model.GetIdOrder()
	order_id := orders.ID
	for i := 0; i < len(r.Form); i++ {

		productID := r.FormValue("productID[" + strconv.Itoa(i) + "]")
		productIDs, _ := strconv.Atoi(productID)
		quantity := r.FormValue("quantity[" + strconv.Itoa(i) + "]")
		quantitys, _ := strconv.Atoi(quantity)
		price := r.FormValue("unit-price[" + strconv.Itoa(i) + "]")
		prices, _ := strconv.Atoi(price)
		if prices != 0 && productIDs != 0 && quantitys != 0 {
			model.SaveOrderItem(quantitys, prices, productIDs, order_id)
		}
	}
	ids := strconv.Itoa(int(order_id))

	http.Redirect(w, r, URL_CHECKOUT+"/"+ids, http.StatusMovedPermanently)

}
