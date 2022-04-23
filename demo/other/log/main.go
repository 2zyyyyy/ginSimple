package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	gin.DisableConsoleColor()

	// logging to file
	file, _ := os.Create("./ginLog.log")
	//gin.DefaultWriter = io.MultiWriter()
	// 同时将日志文件写入文件和控制台
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "gin日志文件生成测试"})
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
