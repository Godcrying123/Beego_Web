package controllers

import (
	"beeblog/models"
	"strconv"

	//"fmt"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	this.TplName = "category.html"
	var controlusername string
	controlusername, this.Data["IsLogin"], this.Data["IsAdmin"] = checkAccount(this.Ctx)
	this.Data["username"] = controlusername
	//beego.Info(isLogin)
	this.Data["Input_Value"] = ""
	var err error
	this.Data["categories"], err = models.GetAllCategory()
	if err != nil {
		beego.Error(err)
		return
	}
	cid, err1 := strconv.ParseInt(this.Input().Get("cid"), 10, 64)
	beego.Info(cid)
	if err1 != nil {
		beego.Error(err1)
		return
	}
	if &cid != nil {
		category, err1 := models.GetCategoryById(cid)
		if err1 != nil {
			beego.Info(err)
			return
		}
		this.Data["Input_Value"] = category.Title
		this.Data["Hidden_Id"] = category.Id
		return
	}
}

func (this *CategoryController) Post() {
	this.TplName = "category.html"
	cbuttonAdd := this.Input().Get("addcate")
	cbuttonUp := this.Input().Get("upcate")
	cateTitle := this.Input().Get("inputcategory")
	cid, err1 := strconv.ParseInt(this.Input().Get("hiddenid"), 10, 64)
	if err1 != nil {
		beego.Error(err1)
	}
	if cbuttonAdd != "" {
		err := models.AddCategory(cateTitle)
		if err != nil {
			beego.Error(err)
			this.Redirect("/category", 302)
			return
		}
		this.Redirect("/category", 302)
		return
	} else if cbuttonUp != "" {
		beego.Info("Title Updated Here")
		err := models.UpdateCategoryById(cid, cateTitle)
		if err != nil {
			beego.Error(err)
			//this.Redirect("/category", 302)
			return
		}
		this.Data["categories"], _ = models.GetAllCategory()
		this.Redirect("/category", 302)
		return
	}
}

func (this *CategoryController) Delete() {
	this.TplName = "category.html"
	cdid, err1 := strconv.ParseInt(this.Input().Get("cdid"), 10, 64)
	if err1 != nil {
		beego.Error(err1)
		return
	}
	err1 = models.DeleteCategoryById(cdid)
	if err1 != nil {
		beego.Error(err1)
		return
	}
	this.Data["categories"], _ = models.GetAllCategory()
	this.Redirect("/category", 302)
	return
}
