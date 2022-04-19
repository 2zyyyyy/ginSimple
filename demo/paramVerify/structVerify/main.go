package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 结构体验证

type Person struct {
	// 不能为空并且大于10
	Age      int       `form:"age" binding:"required,gt=10"`
	Name     string    `form:"name" binding:"required"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	r := gin.Default()
	r.GET("/2zyyyyy", func(ctx *gin.Context) {
		var person Person
		if err := ctx.ShouldBind(&person); err != nil {
			ctx.JSON(500, gin.H{"msg": err})
			return
		}
		ctx.JSON(200, fmt.Sprintf("%#v\n", person))
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
