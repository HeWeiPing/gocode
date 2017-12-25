package main

import (
	"blogProject/controllers"
	"blogProject/models"
	"clog"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

func main() {
	clog.Clogv(clog.Yellow, "main() Start!")
	orm.Debug = true
	orm.RunSyncdb("default", false, true)

	//注册beego路由
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Run()
}
