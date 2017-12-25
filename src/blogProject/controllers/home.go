package controllers

import (
	"clog"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	clog.Clogv(clog.Yellow, "Rq URI=%s", this.Ctx.Input.URI())
	clog.Clogv(clog.Pink, "Rq(exit)=%s", this.Input().Get("exit"))
	clog.Clogv(clog.Yellow, "isExit=%t", this.Input().Get("exit") == "true")
	clog.Clogv(clog.Pink, "Rq uname=%s", this.Ctx.Input.Cookie("uname"))
	clog.Clogv(clog.Pink, "Rq pwd=%s", this.Ctx.Input.Cookie("pwd"))
	clog.Clogv(clog.Pink, "Rq autoLogin=%s", this.Input().Get("autoLogin"))

	this.TplName = "home.html"
	this.Data["IsHome"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	clog.Clogv(clog.Red, "IsLogin=%t", checkAccount(this.Ctx))

}