package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB       *gorm.DB
	REDIS    *redis.Client
	APPYML   *viper.Viper //app的yml配置文件解析器
	DBYML    *viper.Viper
	REDISYML *viper.Viper
	LOG      *zap.Logger
)
