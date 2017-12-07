package controllers

import (
	"github.com/astaxie/beego"
)

// OrderController operations for Order
type OrderController struct {
	beego.Controller
}

func (this *OrderController) Order() {
	this.InitFrontEndTemplate("frontend/order.tpl")
}

func (this *OrderController) InitFrontEndTemplate(TplName string) {
	this.Layout = "frontend/base/layout.tpl"
	this.TplName = TplName
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "frontend/base/header.tpl"
	this.LayoutSections["Footer"] = "frontend/base/footer.tpl"
}
