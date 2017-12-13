package model

import (
	"log"
)

type Category struct {
	Model
	Name     string    `schema:"name"`
	Products []Product // Category has many product
}

func GetAllCategory() (AllCategorys []Category, err error) {
	rows, err := DBCon.Query("SELECT name FROM categories")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		AllCategory := Category{}
		err = rows.Scan(&AllCategory.Name)
		if err != nil {
			return
		}
		AllCategorys = append(AllCategorys, AllCategory)
	}
	rows.Close()
	return AllCategorys, err
}
