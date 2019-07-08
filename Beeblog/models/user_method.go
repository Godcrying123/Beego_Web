package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func AddUser(username, email, password string) error {
	o := orm.NewOrm()
	profile := &Profile{
		FansCount:    0,
		ArticleCount: 0,
	}
	user := &User{
		Username: username,
		Email:    email,
		Password: password,
		UserType: "default",
		Created:  time.Now(),
		Profile:  profile,
	}
	qs := o.QueryTable("user")
	err := qs.Filter("Username", username).One(user)
	if err == nil {
		beego.Info("there is a user with same user name here")
		return nil
	}
	beego.Error(err)
	_, err = o.Insert(user)
	if err != nil {
		beego.Error(err)
		return err
	}
	return nil
}

func GetUserByUserName(username string) (user *User, err error) {
	o := orm.NewOrm()
	user = &User{
		Username: username,
	}
	qs := o.QueryTable("user")
	err = qs.Filter("Username", username).One(user)
	if err == nil {
		return user, err
	} else {
		beego.Error(err)
		return nil, err
	}
}
