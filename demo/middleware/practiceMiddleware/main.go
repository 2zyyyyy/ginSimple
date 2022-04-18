package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 中间件
func myTime(ctx *gin.Context) {
	start := time.Now()
	ctx.Next()
	// 统计时间
	since := time.Since(start)
	fmt.Println("累计用时：", since)
}

func shopIndexHandler(ctx *gin.Context) {
	time.Sleep(5 * time.Second)
	ctx.String(http.StatusOK, "shopIndexHandler~~~")
}

func shopHomeHandler(ctx *gin.Context) {
	time.Sleep(3 * time.Second)
	ctx.String(http.StatusOK, "shopHomeHandler~~~")
}

func main() {
	r := gin.Default()
	// 注册中间件
	r.Use(myTime)
	shoppingGroup := r.Group("shopping")
	{
		shoppingGroup.GET("/index", shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
