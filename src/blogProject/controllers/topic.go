package controllers

import (
	"blogProject/models"
	//	"clog"
	"github.com/astaxie/beego"
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

	title := this.Input().Get("title")
	content := this.Input().Get("content")

	var err error
	err = models.AddTopic(title, content)
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)
	return
}

func (this *TopicController) Add() {
	this.TplName = "topic_add.html"
}
