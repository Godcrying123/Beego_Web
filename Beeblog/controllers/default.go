package controllers

import (
	"beeblog/models"
	_ "beeblog/models"

	//"os"
	"strconv"
	//"time"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplName = "index.html"
	var controlusername string
	controlusername, this.Data["IsLogin"], this.Data["IsAdmin"] = checkAccount(this.Ctx)
	this.Data["username"] = controlusername
	// this.Data["info"] = ""
	// this.Data["success"] = ""
	// this.Data["danger"] = ""
	// this.Data["warning"] = ""
	if this.Input().Get("exit") == "true" {
		this.Ctx.SetCookie("unmae", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
		this.Redirect("/", 302)
		return
	}
	categories, err1 := models.GetAllCategory()
	if err1 != nil {
		beego.Error(err1)
	}
	this.Data["categories"] = categories
	var cidselect = this.Input().Get("cidselect")
	//beego.Info(cidselect)
	if cidselect == "-1" {
		topics, err := models.GetAllTopic()
		if err != nil {
			beego.Error(err)
		}
		this.Data["topics"] = topics
	} else {
		cid, err2 := strconv.ParseInt(cidselect, 10, 64)
		if err2 != nil {
			beego.Error(err2)
		}
		topics, err := models.GetAllTopicByCategory(cid)
		if err != nil {
			beego.Error(err)
			return
		}
		this.Data["topics"] = topics
	}
}

func (this *MainController) Post() {
	this.TplName = "index.html"
	var username = this.Input().Get("username")
	var pwd = this.Input().Get("password")
	var autologin = this.Input().Get("autoLogin") == "on"
	setcookies(this.Ctx, username, pwd, autologin)
	this.Redirect("/", 302)
	return
}
