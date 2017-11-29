package main

import (
    "goweb1/database"
	"goweb1/model"
	"goweb1/config"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	config.SetupEnv()
	db := database.ConnectDB()

	//create table

	db.AutoMigrate(
		model.User{}, 
		model.Product{}, 
		model.Order{},
		model.OrderItem{},
		model.Category{},
		model.Payment{},
	)

	//Add Foreign Key
	db.Model(model.Product{}).AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")
	db.Model(model.Order{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(model.OrderItem{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")
	db.Model(model.OrderItem{}).AddForeignKey("order_id", "orders(id)", "RESTRICT", "RESTRICT")
	db.Model(model.Payment{}).AddForeignKey("order_id", "orders(id)", "RESTRICT", "RESTRICT")

	//Add index
	db.Model(model.Product{}).AddIndex("idx_category_id", "category_id")
	db.Model(model.Order{}).AddIndex("idx_user_id", "user_id")
	db.Model(model.OrderItem{}).AddIndex("idx_product_id", "product_id")
	db.Model(model.OrderItem{}).AddIndex("idx_order_id", "order_id")
	db.Model(model.Payment{}).AddIndex("idx_order_id", "order_id")

	//Add index unique
	db.Model(model.User{}).AddUniqueIndex("idx_user_name", "username")
	db.Model(model.User{}).AddUniqueIndex("idx_email", "email")
}