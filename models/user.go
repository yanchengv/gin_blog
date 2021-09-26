package models

import "gin_blog/lib/global"

type User struct {
	global.INIT_COLUMNS
	Username string `json:"username" gorm:"comment:用户名"`
	Password string `json:"-" gorm:"comment:用户登录密码"`
	Nickname string `json:"nickname" gorm:"default:系统用户;comment:用户昵称"`
	Email    string `json:"email" gorm:"comment: 邮箱"`
	Phone    string `json:"phone" gorm:"comment:手机号"`
}
