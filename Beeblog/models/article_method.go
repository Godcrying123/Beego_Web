package models

import (
	"time"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

func AddTopic(title, content string, category *Category, user *User) error {
	o := orm.NewOrm()
	//beego.Info("Roger that! I am here")
	topic := &Topic{
		Title:     title,
		Content:   content,
		Created:   time.Now(),
		Updated:   time.Now(),
		ReplyTime: time.Now(),
		Category:  category,
		User:      user,
		Author:    user.Username,
	}
	qs := o.QueryTable("topic")
	err := qs.Filter("title", title).One(topic)
	if err == nil {
		beego.Info("Roger that! I am here")
		return nil
	}
	beego.Error(err)
	_, err = o.Insert(topic)
	if err != nil {
		beego.Error(err)
		return err
	}
	return nil
}

func SaveTopic(title, content string, aid int64) error {
	o := orm.NewOrm()
	beego.Info(aid)
	if aid == 0 {
		topic := &Topic{
			Title:     title,
			Content:   content,
			Created:   time.Now(),
			Updated:   time.Now(),
			ReplyTime: time.Now(),
		}
		qs := o.QueryTable("topic")
		err := qs.Filter("title", title).One(topic)
		if err == nil {
			return nil
		}
		_, err = o.Insert(topic)
		if err != nil {
			return err
		}
		return nil
	} else {
		topic := &Topic{Id: aid}
		if o.Read(topic) == nil {
			topic.Title = title
			topic.Content = content
			topic.Updated = time.Now()
			topic.Views++
			if _, err := o.Update(topic); err == nil {
				return nil
			} else {
				beego.Error(err)
				return err
			}
			return nil
		} else {
			beego.Error("read Data cause the error")
			return nil
		}
	}
	return nil
}

func GetAllTopicByUser(user *User) (topic []*Topic, err error) {
	o := orm.NewOrm()
	topic = make([]*Topic, 0)
	_, err = o.QueryTable("topic").Filter("user_id", user.Id).All(&topic)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	return topic, nil
}

func GetAllTopicByCategory(cid int64) (topic []*Topic, err error) {
	o := orm.NewOrm()
	topic = make([]*Topic, 0)
	_, err = o.QueryTable("topic").Filter("category_id", cid).All(&topic)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	return topic, nil
}

func GetAllTopic() (topic []*Topic, err error) {
	o := orm.NewOrm()
	topic = make([]*Topic, 0)
	_, err = o.QueryTable("topic").All(&topic)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	return topic, nil
}

func GetTopicById(aid int64) (*Topic, error) {
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err := qs.Filter("id", aid).One(topic)
	if err != nil {
		return nil, err
	}
	return topic, nil
}

func DeleteTopic(aid int64) (err error) {
	o := orm.NewOrm()
	topic := &Topic{Id: aid}
	if o.Read(topic) == nil {
		if _, err := o.Delete(topic); err == nil {
			return nil
		} else {
			beego.Error(err)
			return err
		}
		return nil
	} else {
		beego.Error("read Data cause the error")
		return nil
	}
	return nil

}
