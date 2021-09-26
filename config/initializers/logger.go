package initializers

import (
	"gin_blog/lib/global"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//使用zap,初始化日志
func InitLogger() {

	global.LOG, _ = zap.NewProduction()
	defer global.LOG.Sync()

	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel) //哪种级别的日志将被写入 c  DebugLevel等
	global.LOG = zap.New(core, zap.AddCaller())
	global.LOG.Sugar()

}

//如何写入日志
func getEncoder() zapcore.Encoder {
	//修改时间编码器,在日志文件中使用大写字母记录日志级别
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder //修改时间显示
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	//Encoder:编码器(如何写入日志)。 我们将使用开箱即用的NewJSONEncoder()，并使用预先设置的ProductionEncoderConfig();
	//{"level":"debug","ts":1572160754.994731,"msg":"Trying to hit GET request for www.sogo.com"}
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	//编码器从JSON Encoder更改为普通Encoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

//WriterSyncer ：指定日志将写到哪里去。我们使用zapcore.AddSync()函数并且将打开的文件句柄传进去;
func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./log/development.log") //把日志写入这个文件中
	return zapcore.AddSync(file)
}
