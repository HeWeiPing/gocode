package routers

import (
	"blogProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//注册beego路由
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topicView", &controllers.TopicViewController{})
	beego.Router("/personalPage", &controllers.PersonalPageController{})
	beego.AutoRouter(&controllers.TopicController{}) //自动路由

}
