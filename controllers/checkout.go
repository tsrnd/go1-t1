package controllers

import (
	"github.com/astaxie/beego"
)

// CheckoutController operations for Checkout
type CheckoutController struct {
	beego.Controller
}

func (this *CheckoutController) Checkout() {
	this.InitFrontEndTemplate("frontend/checkout.tpl")
}

func (this *CheckoutController) InitFrontEndTemplate(TplName string) {
	this.Layout = "frontend/base/layout.tpl"
	this.TplName = TplName
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "frontend/base/header.tpl"
	this.LayoutSections["Footer"] = "frontend/base/footer.tpl"
}
