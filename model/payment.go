package model

import "log"

type Payment struct {
	Model
	Method  string `schema:"method"`
	Status  int    `schema:"status"`
	Address string `schema:"address"`
	Total   int64  `schema:"total"`
	Vat     int64  `schema:"vat"`
	OrderId uint64 `schema:"order_id"`
	Orders  []Order
}

func CreatePayment(method string, status int, address string, total int64, vat int64, order_id uint64) (payment Payment, err error) {
	err = DBCon.QueryRow("INSERT INTO payments(method,status,address,total,vat,order_id) VALUES (?,?,?,?,?,?)", method, status, address, total, vat, order_id).Scan(&payment.ID)
	if err != nil {
		log.Fatal(err)
	}
	return payment, err
}
