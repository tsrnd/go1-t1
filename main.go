package main

import (
	"log"
	"net/http"

	"goweb1/config"
)

func main() {
	db := config.DB()
	cache := config.Cache()
	router := config.Router(db, cache)
	port := config.Port()
	if err := http.ListenAndServe(p, r); err != nil {
		log.Fatal(err)
	}
}
