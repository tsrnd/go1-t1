package model

type OrderItem struct {
	Id        int
	Quantity  int
	Price     int
	ProductId int
	OrderId   uint
	Product   Product // OrderItem belong to Product
	Order     Order   // OrderItem belong to Order
}

func GetAllOrderItems() (AllOrderItems []OrderItem, err error) {
	rows, err := DB.Query("select * from order_items")
	if err != nil {
		return
	}
	for rows.Next() {
		orderItem := OrderItem{}
		err = rows.Scan(&orderItem.Id, &orderItem.Quantity, &orderItem.Price, &orderItem.ProductId, &orderItem.OrderId)
		if err != nil {
			return
		}
		AllOrderItems = append(AllOrderItems, orderItem)
	}
	rows.Close()
	return
}

// func GetOrderItemByOrder(order_id int64) (orderItems []OrderItem, err error) {
// 	orderItems = []OrderItem{}
// 	err = DB.Where("order_id = ?", order_id).Find(&orderItems).Error
// 	return orderItems, err
// }

// func SaveOrderItem(quantity int, price int, product_id int, order_id uint) (*OrderItem, error) {
// 	orderItem := OrderItem{Quantity: quantity, Price: price, ProductId: product_id, OrderId: order_id}
// 	res := DB.Create(&orderItem)
// 	return &orderItem, res.Error
// }
