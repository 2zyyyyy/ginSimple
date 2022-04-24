package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 加载静态文件
	r.Static("statics", "./statics")
	// 模板解析
	r.LoadHTMLGlob("templates/**/*")

	// 返回前端模板
	r.GET("/home", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
