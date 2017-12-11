package controllers

import (
	"github.com/astaxie/beego"
)

// CheckoutController operations for Checkout
type CheckoutController struct {
	beego.Controller
}

func (this *CheckoutController) Checkout() {
	InitFrontEndTemplate(&this.Controller, "frontend/checkout.tpl")
}
