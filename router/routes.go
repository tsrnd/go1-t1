package router
import (

    "github.com/julienschmidt/httprouter"
    ctrl "goweb1/controllers"
    "net/http"
)


func AllRoute() *httprouter.Router {
	return routes()
}

func routes() *httprouter.Router {
    var uc ctrl.UserController
    var hc ctrl.HomeController

    r := httprouter.New()
    r.ServeFiles("/public/*filepath", http.Dir("public"))
    r.GET("/user", uc.GetUser)
    r.GET("/", hc.Home)

    return r
}