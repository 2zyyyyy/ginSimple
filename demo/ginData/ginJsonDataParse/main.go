package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Json 数据解析和绑定

// Login 定义接收数据的结构体
type Login struct {
	// binding:"required" 修饰的字段 若接收为空值 则报错 是必须字段
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	// JSON 绑定
	r.POST("loginJson", func(context *gin.Context) {
		// 声明接收的变量
		var json Login
		// 将request的body中的数据，自动按照JSON格式解析到结构体
		if err := context.ShouldBindJSON(&json); err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据的工具
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if json.User != "admin" || json.Password != "123456" {
			context.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
