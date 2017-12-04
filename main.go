
package main

import (
    "net/http"
    "goweb1/router"
  _ "github.com/go-sql-driver/mysql"
    "goweb1/database"
    "goweb1/model"
    "goweb1/config"
    "os"
)


func main() {
    r := router.Routes()
    http.ListenAndServe(os.Getenv("SERVER_PORT"), r)
}

func init() {
    config.SetupEnv()
    db := database.ConnectDB()
    model.SetDatabase(db)  
    http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
}

