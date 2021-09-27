package controlleradmin

import (
	"gin_blog/lib/global"
	"gin_blog/models"
	"gin_blog/models/scope"

	"github.com/gin-gonic/gin"
)

//获取文章标签列表

func GetTags(c *gin.Context) {
	var tags []models.Tag
	global.DB.Scopes(scope.Paginate(c)).Order("id asc").Find(&tags)
	c.JSON(200, gin.H{
		"status": 200,
		"msg":    "成功",
		"data":   tags,
	})
}

//新增文章标签
//method post
//@Param {name: xx}
func AddTag(c *gin.Context) {
	var tag models.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(200, gin.H{"status": 500, "msg": "参数异常", "data": gin.H{}})
		return
	}
	if tag.Name == "" {
		c.JSON(200, gin.H{"status": 500, "msg": "名称不能为空", "data": gin.H{}})
		c.Abort()
		return
	}
	global.DB.Create(&tag)
	c.JSON(200, gin.H{
		"status": 200,
		"msg":    "创建标签成功",
		"data":   gin.H{"id": tag.ID, "name": tag.Name},
	})

}

//修改文章标签
//method post
//params {id: xx,name: xx}
func UpdateTag(c *gin.Context) {
	var tag models.Tag
	var tag1 models.Tag
	_ = c.ShouldBindJSON(&tag1)

	tagMap := map[string]interface{}{
		"Name": tag1.Name,
	}
	r := global.DB.Where("id = ?", tag1.ID).First(&tag)
	err := r.Updates(tagMap).Error
	if err != nil {
		c.JSON(200, gin.H{"status": 500, "msg": "修改失败", "data": gin.H{}})
		return
	}
	c.JSON(200, gin.H{"status": 200, "msg": "修改成功", "data": gin.H{}})
}

//删除文章标签

func DeleteTag(c *gin.Context) {
	var tag models.Tag
	_ = c.ShouldBindJSON(&tag)
	global.DB.Delete(&models.Tag{}, tag.ID)
	c.JSON(200, gin.H{"status": 200, "msg": "删除成功", "data": gin.H{}})
}
