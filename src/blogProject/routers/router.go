package routers

import (
	"blogProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/", &controllers.LoginController{})
}
