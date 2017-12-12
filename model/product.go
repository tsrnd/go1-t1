package model

type Product struct {
	Id         int
	Name       string
	Price      int
	Image      string
	Size       string
	Color      string
	CategoryId uint
	Category   Category    // Product belong to Category
	OrderItems []OrderItem // Product has many OrderItem

}

func GetAllProduct() (AllProducts []Product, erro error) {
	rows, err := DB.Query("select * from products")
	if err != nil {
		return
	}
	for rows.Next() {
		item := Product{}
		err = rows.Scan(&item.Id, &item.Name, &item.Price, &item.Image, &item.Size, &item.Color)
		if err != nil {
			return
		}
		AllProducts = append(AllProducts, item)
	}
	rows.Close()
	return

}

// func GetProductByCategory(category_id uint, ID int64) (products []Product, erro error) {
// 	products = []Product{}
// 	erro = DB.Where("category_id = ?&& id != ?", category_id, ID).Find(&products).Error
// 	return products, erro
// }

// func GetProductByID(ID int64) (product []Product, erro error) {
// 	product = []Product{}
// 	erro = DB.First(&product, ID).Error
// 	return product, erro
// }
