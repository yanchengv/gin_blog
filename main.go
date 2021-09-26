package main

import (
	"gin_blog/config/initializers"
	"gin_blog/lib/global"
	"gin_blog/routers"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化数据
	initializers.InitData()

	r := gin.Default()
	//基于zap的中间件,使用zap日志库来接收gin框架默认输出的日志。
	r.Use(ginzap.Ginzap(global.LOG, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(global.LOG, true))

	//初始化路由
	routers.InitAdminRouter(r)
	routers.InitApiRouter(r)

	r.Run()

}
