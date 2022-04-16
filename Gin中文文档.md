# Gin中文文档

### 简介

#### 介绍

- Gin是一个golang的微框架，封装比较优雅，API友好，源码注释比较明确，具有快速灵活，容错方便等特点
- 对于golang而言，web框架的依赖要远比Python，Java之类的要小。自身的`net/http`足够简单，性能也非常不错
- 借助框架开发，不仅可以省去很多常用的封装带来的时间，也有助于团队的编码风格和形成规范

#### 安装

要安装Gin软件包，您需要安装Go并首先设置Go工作区。

1.首先需要安装Go（需要1.10+版本），然后可以使用下面的Go命令安装Gin。

> go get -u github.com/gin-gonic/gin

2.将其导入您的代码中：

> import “github.com/gin-gonic/gin”

3.（可选）导入net/http。例如，如果使用常量，则需要这样做http.StatusOK。

> import “net/http”

#### hello world

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建路由
	r := gin.Default()
	// 绑定路由规则 执行的函数
	// gin.Context封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!")
	})
	// 监听端口，默认8080
	// run(里面不指定默认8080)
	err := r.Run(":8000")
	if err != nil {
		return
	}
}
```



### gin 路由

#### 基本路由

- gin 框架中采用的路由库是基于httprouter做的

- 地址为：https://github.com/julienschmidt/httprouter

  

#### Restful风格的API

- gin支持Restful风格的API
- 即Representational State Transfer的缩写。直接翻译的意思是”表现层状态转化”，是一种互联网应用程序的API设计理念：URL定位资源，用HTTP描述操作

1.获取文章 /blog/getXxx Get blog/Xxx

2.添加 /blog/addXxx POST blog/Xxx

3.修改 /blog/updateXxx PUT blog/Xxx

4.删除 /blog/delXxxx DELETE blog/Xxx

#### API参数

- 可以通过Context的Param方法来获取API参数

- localhost:8000/xxx/param

  ```go
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
  ```

  ![image-20220416191615284](C:\Users\73554\AppData\Roaming\Typora\typora-user-images\image-20220416191615284.png)

#### URL参数

- URL参数可以通过DefaultQuery()或Query()方法获取

- DefaultQuery()若参数不存在，返回默认值，Query()若不存在，返回空串

- API ? name=zs

  ```go
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
  ```

  ![image-20220416192318328](C:\Users\73554\AppData\Roaming\Typora\typora-user-images\image-20220416192318328.png)

  传递参数结果：

  ![image-20220416192508172](C:\Users\73554\AppData\Roaming\Typora\typora-user-images\image-20220416192508172.png)

#### 表单参数

- 表单传输为post请求，http常见的传输格式为四种：

  - application/json
  - application/x-www-form-urlencoded
  - application/xml
  - multipart/form-data

- 表单参数可以通过PostForm()方法获取，该方法默认解析的是x-www-form-urlencoded或from-data格式的参数

  ```go
  package main
  
  import (
  	"fmt"
  	"github.com/gin-gonic/gin"
  	"net/http"
  )
  
  func main() {
  	r := gin.Default()
  	r.POST("/form", func(c *gin.Context) {
  		types := c.DefaultPostForm("type", "post")
  		username := c.PostForm("username")
  		password := c.PostForm("password")
  		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
  	})
  	err := r.Run()
  	if err != nil {
  		return 
  	}
  }
  ```

  ![image-20220416193957144](C:\Users\73554\AppData\Roaming\Typora\typora-user-images\image-20220416193957144.png)

#### 上传单个文件

- multipart/form-data格式用于文件上传

- gin文件上传与原生的net/http方法类似，不同在于gin把原生的request封装到c.Request中
  ```go
  package main
  
  import (
  	"fmt"
  	"github.com/gin-gonic/gin"
  	"net/http"
  )
  
  func main() {
  	r := gin.Default()
  	//限制上传最大尺寸
  	r.MaxMultipartMemory = 8 << 20
  	r.POST("/upload", func(c *gin.Context) {
  		file, err := c.FormFile("file")
  		if file == nil {
  			c.String(500, "未选择任何文件")
  		}
  		if err != nil {
  			return
  		}
  		err = c.SaveUploadedFile(file, file.Filename)
  		if err != nil {
  			fmt.Println("saveUploadedFile failed, err:", err)
  			return
  		}
  		c.String(http.StatusOK, file.Filename)
  	})
  	err := r.Run()
  	if err != nil {
  		return
  	}
  }
  ```

  ![image-20220416195142560](C:\Users\73554\AppData\Roaming\Typora\typora-user-images\image-20220416195142560.png)

#### 上传特定文件

有的用户上传文件需要限制上传文件的类型以及上传文件的大小，但是gin框架暂时没有这些函数(也有可能是我没找到)，因此基于原生的函数写法自己写了一个可以限制大小以及文件类型的上传函数

```go
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
```

#### 上传多个文件

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    <form action="http://localhost:8000/upload" method="post" enctype="multipart/form-data">
          上传文件:<input type="file" name="files" multiple>
          <input type="submit" value="提交">
    </form>
</body>
</html>
```

```go
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
```

![image-20220416204148895](C:\Users\73554\AppData\Roaming\Typora\typora-user-images\image-20220416204148895.png)

#### routes group

- routes group是为了管理一些相同的URL

  ```go
  package main
  
  import (
  	"fmt"
  	"github.com/gin-gonic/gin"
  )
  
  func main() {
  	r := gin.Default()
  	// 路由组1 处理get请求
  	v1 := r.Group("/v1")
  	// {} 是书写规范
  	{
  		v1.GET("/login", login)
  		v1.GET("/submit", submit)
  	}
  	v2 := r.Group("/v2")
  	// {} 是书写规范
  	{
  		v2.POST("/login", login)
  		v2.POST("/submit", submit)
  	}
  	err := r.Run(":8080")
  	if err != nil {
  		return
  	}
  }
  
  func submit(context *gin.Context) {
  	name := context.DefaultQuery("name", "寂幻法师")
  	context.String(200, fmt.Sprintf("hello %s\n", name))
  }
  
  func login(context *gin.Context) {
  	name := context.DefaultQuery("name", "东夷战士")
  	context.String(200, fmt.Sprintf("hello %s\n", name))
  }
  ```

  ![image-20220416212013521](C:\Users\73554\AppData\Roaming\Typora\typora-user-images\image-20220416212013521.png)

#### 404页面设置

```go
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
```

![image-20220416212646597](C:\Users\73554\AppData\Roaming\Typora\typora-user-images\image-20220416212646597.png)

### gin 数据解析与绑定

#### Json 数据解析和绑定





#### 表单数据解析和绑定







#### URL数据解析和绑定















