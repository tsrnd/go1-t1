
package main

import (
    "net/http"
    "goweb1/router"
  _ "github.com/go-sql-driver/mysql"
    "goweb1/database"
    "goweb1/model"
)

func main() {
    r := router.AllRoute()
    http.ListenAndServe(":8080", r)
}

func init() {
    db := database.ConnectDB()
    model.SetDatabase(db)
    http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
}

