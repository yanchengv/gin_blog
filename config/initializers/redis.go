package initializers

import (
	"context"
	"gin_blog/lib/global"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

//初始化Redis
func InitRedis() (err error) {
	global.REDIS = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", //没有密码则设置空字符串
		DB:       0,  //默认用的数据库
		// DialTimeout:  10 * time.Second,
		// ReadTimeout:  30 * time.Second,
		// WriteTimeout: 30 * time.Second,
		// PoolSize:     10,
		// PoolTimeout:  30 * time.Second,
	})

	_, err1 := global.REDIS.Ping(ctx).Result()
	if err1 != nil {
		global.LOG.Error("redis 链接失败")
		return err1
	}
	return nil
}
