package models

import (
	"time"
	//"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int64
	Username string `orm:index`
	Password string
	Email    string   `orm:index`
	Profile  *Profile `orm:"rel(one)"`
	UserType string
	Created  time.Time
	Topic    []*Topic `orm:"reverse(many)"`
}

type Profile struct {
	Id           int64
	Name         string
	Age          int64
	Sex          string
	Location     string `orm:index`
	User         *User  `orm:"reverse(one)"`
	FansCount    int64
	ArticleCount int64 `orm:index`
}

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Updated         time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
	Topic           []*Topic `orm:"reverse(many)"`
}

type Topic struct {
	Id              int64
	Title           string
	Content         string `orm:size(5000)`
	Attachement     string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
	Category        *Category `orm:"rel(fk)"`
	User            *User     `orm:"rel(fk)"`
}
