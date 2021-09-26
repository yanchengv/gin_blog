package controllerapi

import (
	"gin_blog/lib/global"
	"gin_blog/models"

	"github.com/gin-gonic/gin"
)

//获取文章标签列表
func GetTags(c *gin.Context) {
	var tag models.Tag
	
	global.DB.First(&tag)
	c.JSON(200, gin.H{
		"status": 200,
		"msg":    "成功api",
		"data":   gin.H{"id": tag.ID},
	})
}
