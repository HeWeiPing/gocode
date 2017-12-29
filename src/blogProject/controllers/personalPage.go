package controllers

import (
	"blogProject/models"
	//	"clog"
	"github.com/astaxie/beego"
)

type PersonalPageController struct {
	beego.Controller
}

func (this *PersonalPageController) Get() {
	var err error
	this.TplName = "personal_page.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["Topics"], err = models.GetAllTopics()
	if err != nil {
		beego.Error(err)
	}
}
