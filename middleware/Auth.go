package middleware
import (
	"net/http"
	ctrl "goweb1/controllers"
)


func Auth(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if !ctrl.AuthIsLogin(r) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	next(w, r)
}