package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/user", func(context *gin.Context) {
		// 指定默认值
		// http://localhost:8080/user 才会输出默认值
		name := context.DefaultQuery("name", "东夷战士")
		context.String(http.StatusOK, fmt.Sprintf("hello %s", name))
		r.NoRoute(func(context *gin.Context) {
			context.String(http.StatusNotFound, "404 not found! \n Power By 东夷战士!")
		})
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
