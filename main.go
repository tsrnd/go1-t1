package main

import (
	"log"
	"net/http"

	"goweb1/config"
)

func main() {
	db := config.DB()
	//cache := config.Cache()
	router := config.Router(db)
	port := ":8080"
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}
