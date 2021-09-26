package middleware

import (
	"fmt"
	"gin_blog/lib/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	j := &utils.JWT{}
	return func(c *gin.Context) {

		token := c.Request.Header.Get("token")
		fmt.Println(token)
		userInfo, err := j.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的token",
				"data": gin.H{},
			})
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("uid", userInfo.ID)
		c.Set("nickname", userInfo.Nickname)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
