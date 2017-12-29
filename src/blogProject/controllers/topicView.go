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

var viewTopicId string

func (this *TopicViewController) Get() {
	var err error
	this.TplName = "topic_view.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	viewTopicId = this.Input().Get("tid")
	clog.Clogv(clog.Pink, "view topic viewTopicId=%s", viewTopicId)

	//get the #id topic
	this.Data["TitleView"], err = models.GetTopicById(viewTopicId)
	if err != nil {
		beego.Error(err)
	}
	//get the #id topic all of replys
	this.Data["Replys"], err = models.GetAllReplysById(viewTopicId)
	if err != nil {
		beego.Error(err)
	}
}

func (this *TopicViewController) Post() {
	var err error
	repl := new(models.Respone)
	op := this.Input().Get("op")
	rp_tid := this.Input().Get("tid")
	rp_name := this.Input().Get("rp_name")
	rp_content := this.Input().Get("rp_content")
	rp_title := this.Input().Get("rp_title")

	repl.TopicId, err = strconv.ParseInt(rp_tid, 10, 64)
	repl.RpName = rp_name
	repl.Title = rp_title
	repl.Content = rp_content

	err = models.ResponeOps(op, repl)
	if err != nil {
		beego.Error(err)
	}

	//increment topic ReplyCount(int64)
	var tmp *models.Topic
	tmp, err = models.GetTopicById(viewTopicId)
	if err != nil {
		beego.Error(err)
	}
	tmp.ReplyCount++
	err = models.TopicOps("mod", tmp)
	if err != nil {
		beego.Error(err)
	}

	clog.Clogv(clog.Yellow, "op=%s", op)
	clog.Clogv(clog.Yellow, "rp_tid=%s", rp_tid)
	clog.Clogv(clog.Yellow, "rp_content=%s", rp_content)
	clog.Clogv(clog.Yellow, "rp_title=%s", rp_title)

	this.Redirect("/topicView?tid="+viewTopicId, 302)
	return
}
