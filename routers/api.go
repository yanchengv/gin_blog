package routers

import (
	controllerapi "gin_blog/controllers/api"
	"gin_blog/middleware"

	"github.com/gin-gonic/gin"
)

//前端路由
func InitApiRouter(r *gin.Engine) {

	api := r.Group("/api")
	api.POST("/sendLoginCode", controllerapi.SendLoginCode)
	api.POST("/login", controllerapi.Login)

	api.Use(middleware.JWTAuthMiddleware())
	api.GET("/getTags", controllerapi.GetTags)

}
