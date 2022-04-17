package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

// 多种相应格式
func main() {
	r := gin.Default()
	// 1.json
	r.GET("/someJson", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "someJson", "status": 200})
	})
	// 2.结构体响应
	r.GET("/someStruct", func(context *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "admin"
		msg.Message = "struct message"
		msg.Number = 9527
		context.JSON(200, msg)
	})
	// 3.xml
	r.GET("/someXML", func(context *gin.Context) {
		context.XML(200, gin.H{"message": "xml message"})
	})
	// 4.YAML
	r.GET("/someYAML", func(context *gin.Context) {
		context.YAML(200, gin.H{"message": "YAML message"})
	})
	// 5.protobuf格式 Google开发的高效存储读取的工具
	// 数组？切片？如果在机构建一个传输格式，应该是什么格式
	r.GET("/someProtoBuf", func(context *gin.Context) {
		resp := []int64{int64(1), int64(2)}
		// 定义数据
		label := "label"
		// 传protobuf格式的诗句
		data := &protoexample.Test{
			Label: &label,
			Reps:  resp,
		}
		context.ProtoBuf(200, data)
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
