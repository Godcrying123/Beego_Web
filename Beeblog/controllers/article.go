package controllers

import (
	"beeblog/models"
	"strconv"

	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Get() {
	this.TplName = "article.html"
	var controlusername string
	controlusername, this.Data["IsLogin"],this.Data["IsAdmin"] = checkAccount(this.Ctx)
	categories, err := models.GetAllCategory()
	if err != nil {
		beego.Error(err)
		return
	}
	this.Data["categories"] = categories
	this.Data["username"] = controlusername
	controluser, err1 := models.GetUserByUserName(controlusername)
	if err1 != nil {
		beego.Error(err1)
		this.Redirect("/article", 302)
	}
	this.Data["articles"], err1 = models.GetAllTopicByUser(controluser)
	if err1 != nil {
		beego.Error(err1)
		this.Redirect("/article", 302)
		return
	}
	aid, err2 := strconv.ParseInt(this.Input().Get("aid"), 10, 64)
	if err2 != nil {
		beego.Error(err2)
		//return
	}
	if &aid != nil {
		topic, err1 := models.GetTopicById(aid)
		if err1 != nil {
			beego.Error(err1)
			return
		}
		category, err2 := models.GetCategoryById(topic.Category.Id)
		if err2 != nil {
			beego.Error(err2)
			return
		}
		this.Data["Input_Title"] = topic.Title
		this.Data["Input_Content"] = topic.Content
		this.Data["Hidden_Id"] = topic.Id
		this.Data["category_id"] = category.Id
		return
	}
}

func (this *ArticleController) Post() {
	this.TplName = "article.html"
	articletitle := this.Input().Get("article_title")
	articlecontent := this.Input().Get("article_content")
	categoryname := this.Input().Get("categoryname")
	//beego.Info(categoryname)
	abtncreate := this.Input().Get("create")
	abtnsave := this.Input().Get("save")
	abtnselect := this.Input().Get("selected")
	username := this.Input().Get("usernamebtn")
	//beego.Info(username)
	//beego.Info(abtnselect)
	aid, err1 := strconv.ParseInt(this.Input().Get("hidden_aid"), 10, 64)
	if err1 != nil {
		beego.Error(err1)
	}
	cid, err2 := strconv.ParseInt(categoryname, 10, 64)
	if err2 != nil {
		beego.Error(err2)
	}
	category, err3 := models.GetCategoryById(cid)
	if err3 != nil {
		beego.Error(err3)
	}
	user, err4 := models.GetUserByUserName(username)
	if err4 != nil {
		beego.Error(err4)
	}
	//beego.Info(aid)
	//beego.Info(articlecontent)
	if abtncreate != "" {
		beego.Info("create the content")
		err := models.AddTopic(articletitle, articlecontent, category, user)
		if err != nil {
			beego.Error(err)
			this.Redirect("/article", 302)
			return
		}
		this.Redirect("/article", 302)
		return
	} else if abtnsave != "" {
		beego.Info("save the content")
		err := models.SaveTopic(articletitle, articlecontent, aid)
		if err != nil {
			beego.Error(err)
		}
		this.Data["articles"], _ = models.GetAllTopic()
		this.Data["categories"], _ = models.GetAllCategory()
		//this.Redirect("/category", 302)
		return
	} else if abtnselect != "" {
		aid, err2 := strconv.ParseInt(abtnselect, 10, 64)
		if err2 != nil {
			beego.Error(err2)
			return
		}
		err2 = models.DeleteTopic(aid)
		if err2 != nil {
			beego.Error(err2)
			return
		}
		this.Redirect("/article", 302)
	}
}

func (this *ArticleController) Delete() {
	this.TplName = "article.html"
}

func (this *ArticleController) Modify() {
	this.TplName = "article.html"
}
