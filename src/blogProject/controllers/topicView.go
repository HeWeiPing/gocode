package controllers

import (
	"blogProject/models"
	"clog"
	"github.com/astaxie/beego"
	"strconv"
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
	this.Data["TitleView"], err = models.GetTopicById(tid)
	if err != nil {
		beego.Error(err)
	}
	clog.Clogv(clog.Red, "IsLogin=%t", checkAccount(this.Ctx))

}

func (this *TopicViewController) Post() {
	var err error
	repl := new(models.Respone)
	op := this.Input().Get("op")
	rp_tid := this.Input().Get("rp_tid")
	rp_name := this.Input().Get("rp_name")
	rp_content := this.Input().Get("rp_content")
	rp_title := this.Input().Get("rp_title")

	repl.TopicId, err = strconv.ParseInt(rp_tid, 10, 64)
	repl.RpName = rp_name
	repl.Title = rp_title
	repl.Content = rp_content

	err = models.ResponeUpdate(op, repl)
	if err != nil {
		beego.Error(err)
	}
}
