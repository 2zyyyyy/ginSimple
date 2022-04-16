package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 路由组1 处理get请求
	v1 := r.Group("/v1")
	// {} 是书写规范
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}
	v2 := r.Group("/v2")
	// {} 是书写规范
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

func submit(context *gin.Context) {
	name := context.DefaultQuery("name", "寂幻法师")
	context.String(200, fmt.Sprintf("hello %s\n", name))
}

func login(context *gin.Context) {
	name := context.DefaultQuery("name", "东夷战士")
	context.String(200, fmt.Sprintf("hello %s\n", name))
}
