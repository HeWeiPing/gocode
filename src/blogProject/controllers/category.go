package controllers

import (
	"blogProject/models"
	"clog"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	this.TplName = "category.html"
	this.Data["IsCategory"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	op := this.Input().Get("op")
	name := this.Input().Get("catName")
	clog.Clogv(clog.Yellow, "op=%s", op)
	clog.Clogv(clog.Yellow, "catName=%s", name)

	var err error
	pcate := &models.Category{
		Title: name,
	}
	err = models.CategoryOps(op, pcate)
	if err != nil {
		beego.Error(err)
	}

	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}
