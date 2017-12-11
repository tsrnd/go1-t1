package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Product struct {
	Id        int64  `orm:"auto"`
	Name      string `orm:"size(128)"`
	Price     int64
	Image     string       `orm:"size(128)"`
	Size      string       `orm:"size(128)"`
	Color     string       `orm:"size(128)"`
	Category  *Category    `orm:"rel(fk)"`
	OrderItem []*OrderItem `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Product))
}

func (m *Product) TableName() string {
	return "products"
}

// AddProduct insert a new Product into database and returns
// last inserted Id on success.
func AddProduct(m *Product) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProductById retrieves Product by Id. Returns error if
// Id doesn't exist
func GetProductById(id int64) (v *Product, err error) {
	o := orm.NewOrm()
	v = &Product{Id: id}
	if err = o.QueryTable(new(Product)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllProducts
func GetAllProducts() (products []*Product, err error) {
	o := orm.NewOrm()
	var articles []*Product
	num, err := o.QueryTable("products").All(&articles)
	fmt.Println(num)
	return articles, err
}
