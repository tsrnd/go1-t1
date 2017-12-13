package model

import "log"

type OrderItem struct {
	Model
	Quantity  int     `schema:"not null"`
	Price     int     `schema:"price"`
	ProductId int     `schema:"product_id"`
	OrderId   uint    `schema:"order_id"`
	Product   Product // OrderItem belong to Product
	Order     Order   // OrderItem belong to Order
}

func GetAllOrderItems() (AllOrderItems []OrderItem, err error) {
	rows, err := DBCon.Query("select quantity, price, product_id, order_id from order_items")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		AllOrderItem := OrderItem{}
		err = rows.Scan(&AllOrderItem.Quantity, &AllOrderItem.Price, &AllOrderItem.ProductId, &AllOrderItem.OrderId)
		if err != nil {
			return
		}
		AllOrderItems = append(AllOrderItems, AllOrderItem)
	}
	defer rows.Close()
	return AllOrderItems, err

}

func GetOrderItemByOrder(order_id int64) (orderItems OrderItem, err error) {
	rows, err := DBCon.Query("select * from order_items where order_id = ?", order_id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return orderItems, err
}

func SaveOrderItem(quantity int, price int, product_id int, order_id uint) (order_item OrderItem, err error) {
	err = DBCon.QueryRow("INSERT INTO order_items(quantity, price, product_id, order_id) VALUES($1,$2,$3,$4)", quantity, price, product_id, order_id).Scan(&order_item.Quantity, &order_item.Price, &order_item.ProductId, &order_item.OrderId)
	if err != nil {
		log.Fatal(err)
	}
	return order_item, err
}
