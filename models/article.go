package models

import (
	"gin_blog/lib/global"
)

type Article struct {
	global.INIT_COLUMNS
	TagId    uint   `json:"tag_id" gorm:"comment:标题ID"`
	Title    string `json:"title" gorm:"comment: 文章标题"`
	Subtitle string `json:"subtitle" gorm:"comment:文章子标题"`
	Content  string `json:"content" gorm:"comment:文章内容"`
	Status   string `json:"status" gorm:"default:1;comment:文章状态 1=开启,2=关闭"`
}
