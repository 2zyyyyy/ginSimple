package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 重定向

func main() {
	r := gin.Default()
	r.GET("/index", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "http:www.baidu.com")
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
