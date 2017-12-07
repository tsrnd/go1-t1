package controllers

import (
	"github.com/astaxie/beego"
	"goweb1/models"
)

type MainController struct {
	beego.Controller
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
