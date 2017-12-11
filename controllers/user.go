package controllers

import (
	"github.com/astaxie/beego"
)

// UserController operations for User
type UserController struct {
	beego.Controller
}

func (this *UserController) Register() {
	InitFrontEndTemplate(&this.Controller, "frontend/register.tpl")
}

func (this *UserController) Login() {
	InitFrontEndTemplate(&this.Controller, "frontend/login.tpl")
}
