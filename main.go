package main

import (
	"goweb1/config"
	"goweb1/database"
	"goweb1/model"
	"goweb1/router"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/csrf"
)

func main() {
	r := router.Routes()
	http.ListenAndServe(os.Getenv("SERVER_PORT"), csrf.Protect([]byte("32-byte-long-auth-key"), csrf.Secure(false))(r))
}

func init() {
	config.SetupEnv()
	db := database.ConnectDB()
	model.SetDatabase(db)
}
