package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//限制上传最大尺寸 8MB(默认为32MB)
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err:%s", err.Error()))
		}
		// 获取所有图片
		files := form.File["files"]
		// 遍历所有图片
		for _, file := range files {
			// 逐个存
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err%s:", err.Error()))
			}
		}
		// 判断是否没有上传文件
		if form == nil {
			c.String(500, "未选择任何文件")
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded successfully", len(files)))
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
