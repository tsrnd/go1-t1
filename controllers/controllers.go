package controllers

import (
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

const (
	URL_HOME = "/"
	URL_LOGIN = "/login"
)