package initializers

import (
	"gin_blog/lib/global"
	"log"

	"github.com/spf13/viper"
)

//初始化yml配置文件
func InitYML() {
	initDBYml()
	initAppYml()
	initRedisYml()
}

//初始化数据库的YML配置文件
func initDBYml() {
	global.DBYML = viper.New()
	global.DBYML.AddConfigPath("config")   //设置配置文件所在目录
	global.DBYML.SetConfigName("database") //设置配置文件名称
	global.DBYML.SetConfigType("yml")      //设置配置文件格式为YAML
	//viper解析配置文件
	if err := global.DBYML.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

//初始化项目APP的yml配置文件
func initAppYml() {
	global.APPYML = viper.New()
	global.APPYML.AddConfigPath("config") //设置配置文件所在目录
	global.APPYML.SetConfigName("app")    //设置配置文件名称
	global.APPYML.SetConfigType("yml")    //设置配置文件格式为YAML
	//viper解析配置文件
	if err := global.APPYML.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

//初始化项目APP的yml配置文件

func initRedisYml() {
	global.REDISYML = viper.New()
	global.REDISYML.AddConfigPath("config") //设置配置文件所在目录
	global.REDISYML.SetConfigName("redis")  //设置配置文件名称
	global.REDISYML.SetConfigType("yml")    //设置配置文件格式为YAML
	//viper解析配置文件
	if err := global.REDISYML.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}
