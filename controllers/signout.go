package controllers

import (

	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/gorilla/sessions"
)

type (
	SignOutController struct{}
)

func (hc *SignOutController) SignOut(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if session.Values["username"] != nil { 
		var store = sessions.NewCookieStore([]byte("something-very-secret"))
		session, _ := store.Get(r, "session-id")
		delete(session.Values, "username")
		session.Options.MaxAge = -1
		session.Save(r, w)
	}
	http.Redirect(w, r, URL_HOME, http.StatusMovedPermanently)
}
