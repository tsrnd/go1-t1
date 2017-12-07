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
}

func init() {
	for _, route := range routes {
		beego.Router(route.URL, route.Handler, route.mappingMethod...)
	}
}
