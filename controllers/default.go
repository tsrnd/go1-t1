package controllers

import (
	"goweb1/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//test model
	models.FindUser(1)
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (this *MainController) Home() {
	this.InitFrontEndTemplate("frontend/home.tpl")
}

func (this *MainController) InitFrontEndTemplate(TplName string) {
	this.Layout = "frontend/base/layout.tpl"
	this.TplName = TplName
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "frontend/base/header.tpl"
	this.LayoutSections["Footer"] = "frontend/base/footer.tpl"
}
