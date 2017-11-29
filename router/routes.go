package router
import (

    "github.com/julienschmidt/httprouter"
    controllers "goweb1/controllers"
    "net/http"
)

func Routes() *httprouter.Router {
    var uc controllers.UserController
    var home_controller controllers.HomeController
    var product_controller controllers.ProductController
    var order_controller controllers.OrderController
    var login_controller controllers.LoginController
    var register_controller controllers.RegisterController
    var checkout_controller controllers.CheckoutController


    r := httprouter.New()
    r.ServeFiles("/public/*filepath", http.Dir("public"))
    r.GET("/user", uc.GetUser)
    r.GET("/", home_controller.Home)
    r.GET("/single-product", product_controller.Product)
    r.GET("/order", order_controller.Order)
    r.GET("/login", login_controller.Login)
    r.GET("/register", register_controller.Register)
    r.GET("/checkout", checkout_controller.Checkout)

    return r
}