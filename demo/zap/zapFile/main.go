package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
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

	logger := zap.New(core)
	sugarLog = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWrite() zapcore.WriteSyncer {
	file, err := os.Create("./jsonEncoder.log")
	if err != nil {
		fmt.Println("create file failed, err:", err)
	}
	return zapcore.AddSync(file)
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
