package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 服务器要给客户端cookie
	r.GET("/cookie", func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("key_cookie")
		if err != nil {
			cookie = "notSet"
			// 给客户端设置cookie（1.maxAge int,单位s 2.path,cookie所在目录
			//3.domain string,域名 4.secure 是否只能通过HTTPS访问 4.httpOnly bool 是否允许别人通过js获取自己的cookie）
			ctx.SetCookie("key_cookie", "value_cookie", 60,
				"/", "localhost", false, true)
		}
		fmt.Printf("cookie的值为:%s\n", cookie)
		ctx.JSON(http.StatusOK, gin.H{"cookie": cookie})
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
