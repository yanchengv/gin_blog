package controllerapi

import (
	"gin_blog/lib/global"
	"gin_blog/models"
	"gin_blog/models/scope"

	"github.com/gin-gonic/gin"
)

//获取文章标签列表
func GetTags(c *gin.Context) {
	var tags []models.Tag
	//分页标签的数据
	tagList := make([]interface{}, len(tags))
	global.DB.Scopes(scope.Paginate(c)).Order("id asc").Find(&tags)
	for _, tag := range tags {
		tagItem := map[string]interface{}{
			"id":         tag.ID,
			"name":       tag.Name,
			"created_at": tag.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		tagList = append(tagList, tagItem)
	}
	c.JSON(200, gin.H{
		"status": 200,
		"msg":    "获取标签列表数据",
		"data":   tagList,
	})
}
