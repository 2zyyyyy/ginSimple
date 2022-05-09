package main

import (
	"fmt"
	"ginSimple/demo/zap/ginZap/config"
	"ginSimple/demo/zap/ginZap/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
)

func main() {
	// load config from config.json
	if len(os.Args) < 1 {
		return
	}
	if err := config.Init(os.Args[1]); err != nil {
		panic(err)
	}
	// init logger
	if err := logger.InitLogger(config.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	gin.SetMode(config.Conf.Mode)
	r := gin.Default()

	// 注册zap相关中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/hello", func(c *gin.Context) {
		// 假设你有一些数据需要记录到日志中
		var (
			name = "wanli"
			age  = 18
		)
		// 记录日志并使用zap.Xxx(key, val)记录相关字段
		zap.L().Debug("this is hello func", zap.String("user", name), zap.Int("age", age))

		c.String(http.StatusOK, "hello google.com")
	})

	addr := fmt.Sprintf(":%v", config.Conf.Port)
	_ = r.Run(addr)
}
