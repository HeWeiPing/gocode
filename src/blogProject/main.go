package main

import (
	"blogProject/controllers"
	"blogProject/models"
	"clog"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	models.RegisterDB()
}

func main() {
	clog.Clogv(clog.Yellow, "main() Start!")
	// 自动建表
	orm.RunSyncdb("default", false, true)

	//注册beego路由
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.AutoRouter(&controllers.TopicController{}) //自动路由
	beego.Run()
}
