package models

import "time"

// 用户表模型
type User struct {
    Id         int64
    Username   string    `json:"username" orm:"unique;size(15);index"`
    Password   string    `orm:"size(32)"`
    Nickname   string    `json:"nickname" orm:"size(15);index"`
    Email      string    `json:"email" orm:"size(50);index"`
    Lastlogin  time.Time `json:"lastlogin" orm:"auto_now_add;type(datetime);index"`
    Logincount int64     `orm:"index"`
    Lastip     string    `json:"lastip" orm:"size(32);index"`
    Authkey    string    `orm:"size(10)"`
    Active     int8
    Permission string `orm:"size(100);index"`
    Avator     string `json:"avator" orm:"size(10);default(/static/upload/default/user-default-60x60.png)"`
    Upcount    int64
    Post       []*Post     `orm:"reverse(many)"`
    Comments   []*Comments `orm:"reverse(many)"`
}
