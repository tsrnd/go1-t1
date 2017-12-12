package router

import (
	"goweb1/controllers"
	"net/http"
)

var uc controllers.UserController
var home_controller controllers.HomeController
var product_controller controllers.ProductController
var order_controller controllers.OrderController
var login_controller controllers.LoginController
var register_controller controllers.RegisterController
var checkout_controller controllers.CheckoutController
var url_notfound controllers.NotFoundController

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home_controller.Home)
	mux.HandleFunc("/user/:id", uc.GetUser)
	mux.HandleFunc("/single-product/:id", product_controller.Product)
	mux.HandleFunc("/order", order_controller.Order)
	mux.HandleFunc("/order/save", order_controller.SaveOrder)
	mux.HandleFunc("/login", login_controller.Login)
	mux.HandleFunc("/register", register_controller.Register)
	mux.HandleFunc("/logout", login_controller.LogOut)
	mux.HandleFunc("/checkout/:id", checkout_controller.Checkout)
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	return mux

}
