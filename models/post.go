package models

import "time"

type Post struct {
    Id       int64
    User     *User       `orm:"rel(fk);index"`
    Title    string      `orm:"size(100);index"`
    Color    string      `orm:"size(7);index"`
    Urlname  string      `orm:"size(100);index"`
    Urltype  string      `orm:"index"`
    Content  string      `orm:"type(text)"`
    Tags     string      `orm:"size(100);index"`
    Posttime time.Time   `orm:"auto_now_add;type(datetime);index"`
    Views    int64       `orm:"index"`
    Status   int8        `orm:"index"`
    Updated  time.Time   `orm:"auto_now_add;type(datetime);index"`
    Istop    int8        `orm:"index"`
    Cover    string      `orm:"size(70);default(/static/upload/default/blog-default-0.png)"`
    Comments []*Comments `orm:"reverse(many)"`
}
