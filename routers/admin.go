package routers

import (
	controlleradmin "gin_blog/controllers/admin"

	"github.com/gin-gonic/gin"
)

//后台路由
func InitAdminRouter(r *gin.Engine) {
	admin := r.Group("/admin")
	admin.GET("/getTags", controlleradmin.GetTags)
	admin.POST("/addTag", controlleradmin.AddTag)
	admin.POST("/updateTag", controlleradmin.UpdateTag)
	admin.POST("/deleteTag", controlleradmin.DeleteTag)
}
