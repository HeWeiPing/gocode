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
	op := this.Input().Get("op")
	clog.Clogv(clog.Yellow, "op=%s", op)
	clog.Clogv(clog.Yellow, "catName=%s", this.Input().Get("catName"))
	switch op {
	case "add":
		title := this.Input().Get("catName")
		if len(title) == 0 {
			break
		}
		err := models.AddCategory(title)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 302)
		return
	case "del":
		cid := this.Input().Get("cid")
		if len(cid) == 0 {
			break
		}

		err := models.DelCategory(cid)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 302)
		return
	}

	var err error
	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}

	this.TplName = "category.html"
	this.Data["IsCategory"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)
}