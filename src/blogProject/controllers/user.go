package controllers

import (
	"blogProject/models"
	"clog"
	"github.com/astaxie/beego"
	//	"github.com/astaxie/beego/context"
)

//var from string

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	op := this.Input().Get("op")
	from = this.Input().Get("from")
	if op == "add" {
		this.TplName = "register.html"
	}
}

func (this *UserController) Post() {
	op := this.Input().Get("op")
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("cfpwd")
	email := this.Input().Get("email")

	clog.Clogv(clog.Yellow, "Rq URI=%s", this.Ctx.Input.URI())
	clog.Clogv(clog.Pink, "Post Rq from=%s", from)
	clog.Clogv(clog.Pink, "Post Rq uname=%s", uname)
	clog.Clogv(clog.Pink, "Post Rq pwd=%s", pwd)
	clog.Clogv(clog.Pink, "Post Rq email=%s", email)

	userInfo := &models.User{
		UserName:  uname,
		UserPwd:   pwd,
		UserEmail: email,
	}

	this.Data["IsAccountExist"] = models.IsAccountExist(uname)
	clog.Clogv(clog.Pink, "IsAccountExist=%t", this.Data["IsAccountExist"])
	clog.Clogv(clog.Pink, "UserInfo=", *userInfo)
	if this.Data["IsAccountExist"] == true { // can not insert as a new user
		clog.Clogv(clog.Yellow, "User:%s is exist", userInfo.UserName)
		this.Redirect(from, 302)
		return
	}

	this.Data["opResult"] = models.UserOps(op, userInfo)
	this.Redirect(from, 302)
	return
}
