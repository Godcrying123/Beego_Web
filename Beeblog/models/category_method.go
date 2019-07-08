package models

import (
	"time"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

func AddCategory(cate string) error {
	o := orm.NewOrm()
	category := &Category{
		Title:   cate,
		Created: time.Now(),
		Updated: time.Now(),
	}
	qs := o.QueryTable("category")
	err := qs.Filter("title", cate).One(category)
	if err == nil {
		return err
	}
	_, err = o.Insert(category)
	if err != nil {
		return err
	}
	return nil
}

func GetAllCategory() (category []*Category, err error) {
	o := orm.NewOrm()
	category = make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err = qs.All(&category)
	return category, err
}

func GetCategoryById(cid int64) (*Category, error) {
	o := orm.NewOrm()
	category := new(Category)
	qs := o.QueryTable("category")
	err := qs.Filter("id", cid).One(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func UpdateCategoryById(cid int64, title string) (err error) {
	o := orm.NewOrm()
	category := &Category{Id: cid}
	if o.Read(category) == nil {
		category.Title = title
		category.Updated = time.Now()
		category.Views++
		if _, err := o.Update(category); err == nil {
			return err
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

func DeleteCategoryById(cdid int64) error {
	o := orm.NewOrm()
	category := &Category{Id: cdid}
	if o.Read(category) == nil {
		if _, err := o.Delete(category); err == nil {
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
