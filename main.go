
package main

import (
    "net/http"
    "goweb1/router"
  _ "github.com/go-sql-driver/mysql"
    "goweb1/database"
    "goweb1/model"
    "goweb1/config"
    "os"
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
    http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
}

