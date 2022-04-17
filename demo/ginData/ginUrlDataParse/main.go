package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Url 数据解析和绑定

type Login struct {
	// binding:"required" 修饰的字段 若接收为空值 则报错 是必须字段
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/:user/:password", func(context *gin.Context) {
		// 声明接收的变量
		var login Login
		// Bind()默认解析并绑定form格式
		// 根据请求头中的content—type自动推断
		if err := context.ShouldBindUri(&login); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if login.User != "admin" || login.Password != "123456" {
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
