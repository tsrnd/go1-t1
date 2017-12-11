package main

import (
	// "database/sql"
    "goweb1/database"
	"goweb1/config"
	_ "github.com/lib/pq"
	"goweb1/database/migrations"
)



func init() {
	config.SetupEnv()
	DB := database.ConnectDB()
	migrations.SetDatabase(DB)  
}

func main() {
	migrations.Run()
}
