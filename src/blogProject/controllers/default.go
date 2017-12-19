package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type LoginController struct {
	beego.Controller
	Name     string
	Password []byte
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

	c.Data["Html"] = "<div>helo go go go </div>"
}

func (c *LoginController) Get() {
	c.Data["Name"] = "hwp"
	c.Data["Password"] = "123456"
	c.TplName = "login.html"
}
