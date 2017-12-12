package model

type Category struct {
	Id       int
	Name     string
	Products []Product // Category has many product
}

func GetAllCategory() (categories []Category, err error) {
	rows, err := DB.Query("select id, name from categories")
	if err != nil {
		return
	}
	for rows.Next() {
		category := Category{}
		err = rows.Scan(&category.Id, &category.Name)
		if err != nil {
			return
		}
		categories = append(categories, category)
	}
	rows.Close()
	return
}
