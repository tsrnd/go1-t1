package controllers

import (
	"github.com/astaxie/beego"
)

// OrderController operations for Order
type OrderController struct {
	beego.Controller
}

func (this *OrderController) Order() {
	InitFrontEndTemplate(&this.Controller, "frontend/order.tpl")
}
