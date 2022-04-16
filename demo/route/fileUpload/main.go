package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	//限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		_, headers, err := c.Request.FormFile("file")
		if err != nil {
			log.Printf("Error when try to get file:%v\n", err)
		}
		// 判断是否没有上传文件
		if headers == nil {
			c.String(500, "未选择任何文件")
		}
		// headers.Header.Get("Content-Type")获取上传的文件的类型
		if headers.Header.Get("Content-Type") != "image/png" {
			c.String(http.StatusOK, "只支持上传png图片")
			return
		}
		// headers.Size 获取文件大小
		if headers.Size > 1024*1024*2 {
			c.String(http.StatusOK, "文件过大，请重新选择")
			return
		}
		if err != nil {
			return
		}
		err = c.SaveUploadedFile(headers, headers.Filename)
		if err != nil {
			fmt.Println("saveUploadedFile failed, err:", err)
			return
		}
		c.String(http.StatusOK, "上传成功"+headers.Filename)
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
