package model

import "log"

type Order struct {
	Model
	Total      int64 `schema:"total"`
	Status     int   `schema:"status"`
	UserId     uint
	OrderItems []OrderItem // Order hasmany OrderItem
	Payments   []Payment
	User       User // Order belong to User
}

func GetOrderByID(ID uint) (order Order, err error) {
	err = DBCon.QueryRow("select total, status, user_id from orders where id = $1", ID).Scan(&order.ID, &order.Total, &order.UserId)
	if err != nil {
		log.Fatal(err)
	}
	return order, err
}

func SaveOrder(total int64, status int64, user_id uint) (order Order, err error) {
	err = DBCon.QueryRow("INSERT INTO orders(total,status,user_id) VALUES($1,$2,$3) returning id;", total, status, user_id).Scan(&order.ID)
	if err != nil {
		log.Fatal(err)
	}
	return order, err
}

func GetIdOrder() (order Order, err error) {
	err = DBCon.QueryRow("select id from order limit 1").Scan(&order.ID)
	if err != nil {
		log.Fatal(err)
	}
	return order, err
}
