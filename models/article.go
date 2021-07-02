package models

import "github.com/beego/beego/v2/client/orm"

type Article struct {
	Id         int    `form:"-"`
	Title      string `form:"title"`
	Content    string `form:"content"`
	CreateTime int    `orm:"null"`
	//User *User	`orm:"rel(one)"`
}

func (Article) TableName() string {
	return "article"
}

func (Article) DbName() string {
	return "blog"
}

func init() {
	orm.RegisterModel(new(Article))
}
