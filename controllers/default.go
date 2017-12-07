package controllers

import (
	"github.com/astaxie/beego"
	"goweb1/models"
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
