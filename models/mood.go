package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// 心情表
type Mood struct {
	Id       int64
	Content  string    `orm:"type(text)"`
	Cover    string    `orm:"size(70);default(/static/upload/default/blog-default-0.png)"`
	Posttime time.Time `orm:"type(datetime);index"`
}

func (m *Mood) TableName() string {
	return TableName("mood")
}

func (m *Mood) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}
