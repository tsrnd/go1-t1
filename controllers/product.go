package controllers

import (
	"github.com/astaxie/beego"
)

// ProductController operations for Product
type ProductController struct {
	beego.Controller
}

func (this *ProductController) Product() {
	InitFrontEndTemplate(&this.Controller, "frontend/product.tpl")
}
