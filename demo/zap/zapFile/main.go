package main

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
)

// 将日志写入文件

var sugarLog *zap.SugaredLogger

func main() {
	InitLogger()
	sugaredLoggerHttp("www.google3.com")
	sugaredLoggerHttp("https://www.sogo.com")
	defer func(sugarLog *zap.SugaredLogger) {
		_ = sugarLog.Sync()
	}(sugarLog)
}

func InitLogger() {
	// 指定日志将写到哪里去
	writeSyncer := getLogWrite()
	// 编码器(如何写入日志)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller()) // 添加将调用函数信息记录到日志中的功能
	sugarLog = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	// 修改时间编码器  在日志文件中使用大写字母记录日志级别
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	// 将编码器从JSON Encoder更改为普通Encoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWrite() zapcore.WriteSyncer {
	//file, err := os.Create("./consoleEncoder.log")
	//if err != nil {
	//	fmt.Println("create file failed, err:", err)
	//}

	// zap本身不支持文件切割 需要使用Lumberjack进行日志切割归档
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./consoleEncoder.log", // 日志文件位置
		MaxSize:    10,                     // 切割之前 日志文件最大大小（mb）
		MaxAge:     30,                     // 保留历史文件的最大天数
		MaxBackups: 5,                      // 保留历史文件的最大个数
		Compress:   false,                  // 是否压缩/归档历史文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

// zap Sugared Logger
func sugaredLoggerHttp(url string) {
	sugarLog.Debugf("Trying to hit Get request for %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLog.Errorf(
			"Error fetching URL %s : Error = %s",
			url,
			err,
		)
	} else {
		sugarLog.Infof(
			"Success! statusCode = %s for URL %s",
			resp.Status,
			url,
		)
		_ = resp.Body.Close()
	}
}
