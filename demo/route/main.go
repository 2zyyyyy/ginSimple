package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	r := gin.Default()
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		// 截取
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+"is"+action)
	})
	// 监听端口默认为8080
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
