package controllers

import (
	"github.com/astaxie/beego"
)

// ProductController operations for Product
type ProductController struct {
	beego.Controller
}

func (this *ProductController) Product() {
	this.InitFrontEndTemplate("frontend/product.tpl")
}

func (this *ProductController) InitFrontEndTemplate(TplName string) {
	this.Layout = "frontend/base/layout.tpl"
	this.TplName = TplName
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "frontend/base/header.tpl"
	this.LayoutSections["Footer"] = "frontend/base/footer.tpl"
}
