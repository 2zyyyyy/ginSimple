package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		// 指定默认值
		// http://localhost:8080/user才会打印出默认值
		name := c.DefaultQuery("name", "东夷战士")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	// 监听端口默认为8080
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
