package controllers

import (
	"clog"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var from string

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	clog.Clogv(clog.Pink, "Post Rq from=%s", this.Input().Get("from"))
	clog.Clogv(clog.Yellow, "Rq URI=%s", this.Ctx.Input.URI())
	clog.Clogv(clog.Pink, "Rq(exit)=%s", this.Input().Get("exit"))
	clog.Clogv(clog.Yellow, "isExit=%t", this.Input().Get("exit") == "true")
	clog.Clogv(clog.Pink, "Rq uname=%s", this.Ctx.Input.Cookie("uname"))
	clog.Clogv(clog.Pink, "Rq pwd=%s", this.Ctx.Input.Cookie("pwd"))
	clog.Clogv(clog.Pink, "Rq autoLogin=%s", this.Input().Get("autoLogin"))

	from = this.Input().Get("from")
	isExit := this.Input().Get("exit") == "true"
	if isExit {
		this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
		this.Redirect("/", 302)
		clog.Clogv(clog.Green, "收到退出请求，清除cookid")
		return
	}

	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	clog.Clogv(clog.Yellow, "Rq URI=%s", this.Ctx.Input.URI())
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"

	clog.Clogv(clog.Pink, "Post Rq from=%s", from)
	clog.Clogv(clog.Pink, "Post Rq uname=%s", uname)
	clog.Clogv(clog.Pink, "Post Rq pwd=%s", pwd)
	clog.Clogv(clog.Pink, "Post Rq autoLogin=%s", this.Input().Get("autoLogin"))

	if (beego.AppConfig.String("uname") == uname) &&
		(beego.AppConfig.String("pwd") == pwd) {
		maxAge := 0
		if autoLogin {
			maxAge = 600 //600s
		}
		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
		clog.Clogv(clog.Green, "验证通过，保存cookie,重定向到%s", from)
		if from == "topic/add" {
			this.Redirect("/topic/add", 302)
		} else {
			this.Redirect(from, 302)
		}
	} else {
		clog.Clogv(clog.Red, "帐号或密码错误，重定向到/login")
		this.Redirect("/login?from="+from, 302)
	}

	return
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value

	return beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd
}
