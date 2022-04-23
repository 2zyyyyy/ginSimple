package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

// gin 实现热加载

func main() {
	gin.DisableConsoleColor()

	// logging to file
	file, _ := os.Create("./ginLog.log")
	//gin.DefaultWriter = io.MultiWriter()
	// 同时将日志文件写入文件和控制台
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	r := gin.Default()
	r.GET("/gin-realtime-loading", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "使用第三方库air实现热加载", "命令": "air"})
	})
	err := r.Run(":8088")
	if err != nil {
		return
	}
}
