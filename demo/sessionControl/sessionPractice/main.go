package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 模拟实现权限验证中间件
// 有2个路由，login和home
// login用于设置cookie
// home是访问查看信息的请求
// 在请求home之前，先跑中间件代码，检验是否存在cookie
// 访问home，会显示错误，因为权限校验未通过

// 权限校验中间件
func permissionMiddleware() gin.HandlerFunc {
	// 获取客户端cookie并校验
	return func(ctx *gin.Context) {
		if cookie, err := ctx.Cookie("abc"); err == nil {
			if cookie == "123" {
				ctx.Next()
				return
			}
		}
		// 返回错误
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "cookie 鉴权失败"})
		// 若验证不通过，不再调用后续的函数处理
		ctx.Abort()
		return
	}
}

func main() {
	r := gin.Default()
	r.GET("/login", func(ctx *gin.Context) {
		// 设置 cookie(domain 如果设置了 localhost 浏览器不能通过 127.0.0.1 来访问 否则 cookie 会设置失败)
		ctx.SetCookie("abc", "123", 60, "/",
			"localhost", false, true)
		// 返回信息
		ctx.JSON(http.StatusOK, gin.H{"msg": "login success"})
	})
	r.GET("/home", permissionMiddleware(), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "welcome to my home"})
	})
	err := r.Run()
	if err != nil {
		return
	}
}
