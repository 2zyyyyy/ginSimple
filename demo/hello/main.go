package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建路由
	r := gin.Default()
	// 绑定路由规则 执行的函数
	// gin.Context封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!")
	})
	// 监听端口，默认8080
	// run(里面不指定默认8080)
	err := r.Run(":8000")
	if err != nil {
		return
	}
}
