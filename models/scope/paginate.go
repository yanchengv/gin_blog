package scope

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//自定义分页查询
func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageSize := 30
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		offet := (page - 1) * pageSize
		return db.Offset(offet).Limit(pageSize)
	}
}
