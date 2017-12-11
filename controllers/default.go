package controllers

import (
	"fmt"
	"goweb1/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//test model

	articles, _ := models.GetAllProducts()
	c.Data["records"] = articles
	fmt.Println(articles)
	InitFrontEndTemplate(&c.Controller, "frontend/home.tpl")
}

func (this *MainController) Home() {
	InitFrontEndTemplate(&this.Controller, "frontend/home.tpl")
}

func InitFrontEndTemplate(this *beego.Controller, TplName string) {
	this.Layout = "frontend/base/layout.tpl"
	this.TplName = TplName
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "frontend/base/header.tpl"
	this.LayoutSections["Footer"] = "frontend/base/footer.tpl"
}
