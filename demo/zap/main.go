package main

import (
	"go.uber.org/zap"
	"net/http"
)

// 日志库 zap 使用

// 定义全局logger对象
var (
	logger      *zap.Logger
	sugarLogger *zap.SugaredLogger
)

func main() {
	InitLogger()
	defer func(sugarLogger *zap.SugaredLogger) {
		_ = sugarLogger.Sync()
	}(sugarLogger)
	sugaredLoggerHttp("www.google3.com")
	sugaredLoggerHttp("https://www.sogo.com")
}

func InitLogger() {
	logger, _ = zap.NewProduction()
	sugarLogger = logger.Sugar()
}

// zap Logger
func loggerHttp(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url...",
			zap.String("url", url),
			zap.Error(err),
		)
	} else {
		logger.Info(
			"success..",
			zap.String("code", resp.Status),
			zap.String("url", url),
		)
		_ = resp.Body.Close()
	}
}

// zap Sugared Logger
func sugaredLoggerHttp(url string) {
	sugarLogger.Debugf("Trying to hit Get request for %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf(
			"Error fetching URL %s : Error = %s",
			url,
			err,
		)
	} else {
		sugarLogger.Infof(
			"Success! statusCode = %s for URL %s",
			resp.Status,
			url,
		)
		_ = resp.Body.Close()
	}
}
