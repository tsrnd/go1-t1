package routers

import (
	"goweb1/controllers"

	"github.com/astaxie/beego"
)

type Route struct {
	URL           string
	Handler       beego.ControllerInterface
	mappingMethod []string
}

type Routes = []Route

var routes = Routes{
	Route{
		"/",
		&controllers.MainController{},
		[]string{
			"get:Home",
		},
	},
	Route{
		"/login",
		&controllers.UserController{},
		[]string{
			"get:Login",
		},
	},
	Route{
		"/register",
		&controllers.UserController{},
		[]string{
			"get:Register",
		},
	},
	Route{
		"/order",
		&controllers.OrderController{},
		[]string{
			"get:Order",
		},
	},
	Route{
		"/checkout",
		&controllers.CheckoutController{},
		[]string{
			"get:Checkout",
		},
	},
	Route{
		"/product",
		&controllers.ProductController{},
		[]string{
			"get:Product",
		},
	},
}

func init() {
	for _, route := range routes {
		beego.Router(route.URL, route.Handler, route.mappingMethod...)
	}
}
