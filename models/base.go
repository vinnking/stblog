package models

import (
	"fmt"
	"github.com/astaxie/beego"
)

// 返回带前缀的表名
func TableName(str string) string {
	return fmt.Sprintf("%s%s", beego.AppConfig.String("dbprefix"), str)
}
