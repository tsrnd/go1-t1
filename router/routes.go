package router

import (
	controllers "goweb1/controllers"
	middleware "goweb1/middleware"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

var uc controllers.UserController
var home_controller controllers.HomeController
var product_controller controllers.ProductController
var order_controller controllers.OrderController
var login_controller controllers.LoginController
var register_controller controllers.RegisterController
var checkout_controller controllers.CheckoutController
var url_notfound controllers.NotFoundController

func Routes() *httprouter.Router {
    r := httprouter.New()
    r.ServeFiles("/public/*filepath", http.Dir("public"))
    r.GET("/user/:id", uc.GetUser)
    r.POST("/user/:id", uc.UpdateUser)
    r.GET("/", home_controller.Home)
    r.GET("/single-product/:id", product_controller.Product)
    r.GET("/order", order_controller.Order)
    r.POST("/order", order_controller.SaveOrder)
    r.GET("/login", login_controller.Login)
    r.GET("/register", register_controller.Register)
    r.GET("/checkout/:id", checkout_controller.Checkout)
    r.POST("/login", login_controller.ProcessLogin)
    r.POST("/register", register_controller.RegisterPost)
    r.GET("/notfound", url_notfound.NotFound)
  	r.POST("/checkout", checkout_controller.CheckoutPost)

	nAuth := negroni.New()
	nAuth.Use(negroni.HandlerFunc(middleware.Auth))
	nAuth.UseHandlerFunc(login_controller.LogOut)
	r.Handler("GET", "/logout", nAuth)

	return r
}
