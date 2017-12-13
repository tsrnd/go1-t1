package model

import "log"

type Product struct {
	Model
	Name       string      `schema:"name"`
	Price      int         `schema:"price"`
	Image      string      `schema:"image"`
	Size       string      `schema:"size"`
	Color      string      `schema:"color"`
	CategoryId uint        `schema:"category_id"`
	Category   Category    // Product belong to Category
	OrderItems []OrderItem // Product has many OrderItem

}

func GetAllProduct() (AllProducts []Product, err error) {
	rows, err := DBCon.Query("select id, name, price, image, size,  color, category_id from products")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		AllProduct := Product{}
		err = rows.Scan(&AllProduct.ID, &AllProduct.Name, &AllProduct.Price, &AllProduct.Image, &AllProduct.Size, &AllProduct.Color, &AllProduct.CategoryId)
		if err != nil {
			return
		}
		AllProducts = append(AllProducts, AllProduct)
	}
	defer rows.Close()
	return AllProducts, err
}

func GetProductByCategory(category_id uint, ID int64) (products Product, err error) {
	err = DBCon.QueryRow("select name, price, image, size,  color, category_id from products where category_id = $1 and id =$2", category_id, ID).Scan(&products.ID, &products.Name, &products.Price, &products.Image, &products.Size, &products.Color, &products.CategoryId)
	return products, err
}

func GetProductByID(ID int64) (product Product, err error) {
	err = DBCon.QueryRow("select name, price, image, size,  color, category_id  from products where  id =$1", ID).Scan(&product.Name, &product.Price, &product.Image, &product.Size, &product.Color, &product.CategoryId)
	return product, err
}
