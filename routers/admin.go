package routers

import (
	controlleradmin "gin_blog/controllers/admin"

	"github.com/gin-gonic/gin"
)

//后台路由
func InitAdminRouter(r *gin.Engine) {
	admin := r.Group("/admin")
	tag := admin.Group("/tags")
	tag.GET("/", controlleradmin.GetTags)
	tag.POST("/create", controlleradmin.AddTag)
	tag.POST("/update", controlleradmin.UpdateTag)
	tag.POST("/delete", controlleradmin.DeleteTag)

	article := admin.Group("/articles")
	article.GET("/", controlleradmin.GetArticles)
	article.POST("/create", controlleradmin.AddArticle)
	article.POST("/delete", controlleradmin.DeleteArticle)
	article.POST("/update", controlleradmin.UpdateAticle)
}
