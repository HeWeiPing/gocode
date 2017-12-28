package controllers

import (
	"blogProject/models"
	"clog"
	"github.com/astaxie/beego"
	//"strconv"
)

type TopicViewController struct {
	beego.Controller
}

func (this *TopicViewController) Get() {
	clog.Clogv(clog.Yellow, "Rq URI=%s", this.Ctx.Input.URI())
	clog.Clogv(clog.Pink, "Rq(exit)=%s", this.Input().Get("exit"))
	clog.Clogv(clog.Yellow, "isExit=%t", this.Input().Get("exit") == "true")
	clog.Clogv(clog.Pink, "Rq uname=%s", this.Ctx.Input.Cookie("uname"))
	clog.Clogv(clog.Pink, "Rq pwd=%s", this.Ctx.Input.Cookie("pwd"))
	clog.Clogv(clog.Pink, "Rq autoLogin=%s", this.Input().Get("autoLogin"))

	var err error
	this.TplName = "topic_view.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	tid := this.Input().Get("tid")
	clog.Clogv(clog.Pink, "view topic tid=%s", tid)
	//tid, err = strconv.ParseInt(tid, 10, 64)
	this.Data["TitleView"], err = models.GetTopicById(tid)
	if err != nil {
		beego.Error(err)
	}
	clog.Clogv(clog.Red, "IsLogin=%t", checkAccount(this.Ctx))

}
