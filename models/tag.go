package models

import (
	"gin_blog/lib/global"
)

type Tag struct {
	global.INIT_COLUMNS
	Name string `json:"name" gorm:"comment:标签名称"`
}

//如果模型和表面不对应，可以手动指定关联的是数据库那个表
func (Tag) TableName() string {
	return "tags"
}
