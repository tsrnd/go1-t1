
package main

import (
    "net/http"
    "goweb1/router"
  _ "github.com/go-sql-driver/mysql"
)

func main() {
    r := router.AllRoute()
    http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
    http.ListenAndServe(":8080", r)
}
