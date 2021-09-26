package initializers

import "gin_blog/lib/global"

func InitData() {
	InitYML()
	InitDB()
	InitRedis()
	//初始化zap日志
	InitLogger()
	//初始化数据库表
	PostgreTable(global.DB)
}
