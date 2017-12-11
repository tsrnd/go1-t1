package migrations

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func SetDatabase(database *sql.DB) {
	DB = database
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Run() {
	//Create table
	CreateUsersTable()
	CreateCategoriesTable()
	CreateOrderItemsTable()
	CreateOrdersTable()
	CreatePaymentsTable()
	CreateProductsTable()

	//Add foreign keys
	OrderItemsForeignKeysProductId()
	OrderItemsForeignKeysOrderId()
	OrdersForeignKeysUserId()
	PaymentsForeignKeysOrderId()
	ProductsForeignKeysCategoryId()

	//Add index keys
	OrderItemsIndexOrderId()
	OrderItemsIndexProductId()
	OrdersIndexUserId()
	PaymentsIndexOrderId()
	ProductsIndexCategoryId()

	//Add unique keys
	UsersUniqueEmail()
	UsersUniqueUsername()
}