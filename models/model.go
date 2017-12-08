package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

// Init connect dataabse with postgres
func Init() {
	database := beego.AppConfig.String("database")
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	dbsslmode := beego.AppConfig.String("sslmode")
	dbalias := beego.AppConfig.String("dbalias")
	orm.RegisterDriver(database, orm.DRPostgres)
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s password=%s",
		dbhost,
		dbuser,
		dbname,
		dbport,
		dbsslmode,
		dbpassword,
	)
	orm.RegisterDataBase(dbalias, database, dsn)
}
