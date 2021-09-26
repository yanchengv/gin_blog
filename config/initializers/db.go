package initializers

import (
	"gin_blog/lib/global"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//初始化数据库
//var DB *sql.DB

func InitDB() {

	//dsn := "host=localhost user=postgres password=123 dbname=go_mars port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	//dsn := "host=rm-2zefzzv3s8d74dyfvho.pg.rds.aliyuncs.com user=aranya_staging password=e23TsmZasGyEAqd2018 dbname=arnaya_web_dev port=3432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "host=" + global.DBYML.GetString(`database.host`) + " user=" + global.DBYML.GetString(`database.username`) + " password=" + global.DBYML.GetString(`database.password`) + " dbname=" + global.DBYML.GetString(`database.db_name`) + " port=" + global.DBYML.GetString(`database.port`) + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //让gorm输出执行的sql,设置日志级别为Info
	})
	global.DB = db

	if err != nil {
		panic("数据库连接失败")
	}
	sqlDB, _ := global.DB.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
}
