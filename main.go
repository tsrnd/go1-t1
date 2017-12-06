package main

import (
	_ "goweb1/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
orm.RegisterDriver("postgres", orm.DRPostgres)

orm.RegisterDataBase("default", "postgres", "user=postgres password=123456 dbname=goweb sslmode=disable")
}
func main() {
	beego.Run()
}

