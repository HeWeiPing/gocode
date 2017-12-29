package controllers

import (
	"blogProject/models"
	//	"clog"
	"github.com/astaxie/beego"
	"strconv"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	var err error
	this.TplName = "topic.html"
	this.Data["IsTopic"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["Topics"], err = models.GetAllTopics()
	if err != nil {
		beego.Error(err)
	}
}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	var tid int64
	var err error
	op := this.Input().Get("op")
	idstr := this.Input().Get("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	tid, err = strconv.ParseInt(idstr, 10, 64)
	tp := &models.Topic{
		Id:      tid,
		Title:   title,
		Content: content,
	}

	//err = models.TopicUpdate(op, tid, title, content)
	//if err != nil {
	//	beego.Error(err)
	//}

	err = models.TopicOps(op, tp)
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)
	return
}

func (this *TopicController) Add() {
	this.TplName = "topic_add.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)
}

func (this *TopicController) Modify() {
	this.TplName = "topic_mod.html"
}
