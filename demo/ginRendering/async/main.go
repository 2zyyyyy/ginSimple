package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()
	// 1.异步
	r.GET("async", func(context *gin.Context) {
		// 需要一个副本
		copyContext := context.Copy()
		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行：" + copyContext.Request.URL.Path)
		}()
		context.JSON(200, "同步~")
	})
	// 2.同步
	r.GET("sync", func(context *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行：" + context.Request.URL.Path)
		context.JSON(200, "异步")
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
