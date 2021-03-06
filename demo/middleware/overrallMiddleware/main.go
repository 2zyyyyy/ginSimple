package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 全局中间件

// MiddleWare 定义中间件
func MiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行...")
		// 设置变量到 context 的 key 中，可以通过 get()获取
		context.Set("request", "中间件")
		status := context.Writer.Status()
		fmt.Println("中间件执行结束!", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func main() {
	r := gin.Default()
	// 注册中间件
	r.Use(MiddleWare())
	// {} 代码规范
	{
		r.GET("/middleware", func(context *gin.Context) {
			// 取值
			request, _ := context.Get("request")
			fmt.Println("request:", request)
			// 页面接收
			context.JSON(200, gin.H{"request": request})
		})
	}
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
