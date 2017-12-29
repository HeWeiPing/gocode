package main

import (
	"blogProject/models"
	_ "blogProject/routers"
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

	beego.Run()
}
