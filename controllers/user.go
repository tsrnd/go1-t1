package controllers

import (
	"github.com/astaxie/beego"
)

// UserController operations for User
type UserController struct {
	beego.Controller
}

func (this *UserController) Register() {
	this.InitFrontEndTemplate("frontend/register.tpl")
}

func (this *UserController) Login() {
	this.InitFrontEndTemplate("frontend/login.tpl")
}

func (this *UserController) InitFrontEndTemplate(TplName string) {
	this.Layout = "frontend/base/layout.tpl"
	this.TplName = TplName
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "frontend/base/header.tpl"
	this.LayoutSections["Footer"] = "frontend/base/footer.tpl"
}
