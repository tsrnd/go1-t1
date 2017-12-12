package model

type Order struct {
	Id         int
	Total      int64
	Status     int
	UserId     uint
	OrderItems []OrderItem // Order hasmany OrderItem
	Payments   []Payment
	User       User // Order belong to User
}

func GetOrderByID(ID uint) (order Order, err error) {
	order = Order{}
	err = DB.QueryRow("id,total,status from posts where id = $1 order by id desc limit 1", ID).Scan(&order.Id, &order.Total, &order.Status)
	return
}

func SaveOrder(total int64, user_id uint) (err error) {
	order := Order{Total: total, UserId: user_id, Status: 0}
	statement := "insert into posts (total, user_id,status) values ($1, $2, $3)"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(order.Total, order.UserId, order.Status).Scan(&order.Id)
	return
}

func GetIdOrder() (order Order, err error) {
	order = Order{}
	err = DB.QueryRow("id,total from posts order by id desc limit 1").Scan(&order.Id)
	return
}
