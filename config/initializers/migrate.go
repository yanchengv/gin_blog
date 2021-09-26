package initializers

import (
	"gin_blog/lib/global"
	"gin_blog/models"
	"os"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// PostgreTable
//@author: yan
//@function: PostgreTable
//@description: 注册数据库表专用
//@param: db *gorm.DB

func PostgreTable(db *gorm.DB) {
	err := db.AutoMigrate(
		models.Article{},
		models.User{},
		models.Tag{},
	)

	if err != nil {
		global.LOG.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.LOG.Info("register table success")
}
