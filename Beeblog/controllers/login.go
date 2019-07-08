package controllers

import (
	//"context"

	"beeblog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"

}

func (this *LoginController) Post() {
	this.TplName = "login.html"
	uname := this.Input().Get("username")
	email := this.Input().Get("email")
	pwd := this.Input().Get("password")
	autologin := this.Input().Get("autoLogin") == "on"
	// beego.Info(uname)
	// beego.Info(email)
	// beego.Info(pwd)
	// beego.Info(autologin)
	err := models.AddUser(uname, email, pwd)
	if err != nil {
		beego.Error(err)
	}
	setcookies(this.Ctx, uname, pwd, autologin)
	this.Redirect("/", 302)
	return
}

func setcookies(ctx *context.Context, username, password string, autologin bool) {
	maxAge := 0
	if autologin {
		maxAge = 1<<31 - 1
	}
	ctx.SetCookie("uname", username, maxAge, "/")
	ctx.SetCookie("pwd", password, maxAge, "/")
}

func checkAccount(ctx *context.Context) (string, bool, bool) {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return "", false, false
	}
	uname := ck.Value
	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return "", false, false
	}
	pwd := ck.Value
	user, err1 := models.GetUserByUserName(uname)
	if err1 != nil {
		beego.Error(err1)
		return "", false, false
	} else {
		return user.Username, uname == user.Username && pwd == user.Password, user.UserType == "admin"
	}
	//return uname == beego.AppConfig.String("account") && pwd == beego.AppConfig.String("password")
}
