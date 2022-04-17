package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./user/index.html")
	//r.LoadHTMLGlob("htmlTemplate/*")
	r.GET("/index", func(context *gin.Context) {
		context.HTML(200, "user/index.html", gin.H{"title": "万里测试", "address": "www.google.com"})
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
