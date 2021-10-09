package controlleradmin

import (
	"gin_blog/lib/global"
	"gin_blog/models"
	"gin_blog/models/scope"

	"github.com/gin-gonic/gin"
)

//获取文章列表
func GetArticles(c *gin.Context) {
	var articles []models.Article
	articleList := make([]interface{}, 0, len(articles))
	global.DB.Scopes(scope.Paginate(c)).Order("id asc").Find(&articles)
	for _, article := range articles {
		articleItem := map[string]interface{}{
			"id":         article.ID,
			"title":      article.Title,
			"subtitle":   article.Subtitle,
			"content":    article.Content,
			"created_at": article.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		articleList = append(articleList, articleItem)
	}
	c.JSON(200, gin.H{
		"status": 200,
		"msg":    "文章列表",
		"data":   articleList,
	})
}

//创建文章列表
func AddArticle(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(200, gin.H{"status": 200, "msg": "创建失败", "data": gin.H{"error": err}})
		return
	}
	global.DB.Create(&article)
	c.JSON(200, gin.H{
		"status": 200,
		"msg":    "创建成功",
		"data":   gin.H{},
	})
}

//修改文章
func UpdateAticle(c *gin.Context) {
	
}

//删除文章
func DeleteArticle(c *gin.Context) {

}
