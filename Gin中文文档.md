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

- 客户端传参，后端接收并解析到结构体

  ```go
  package main
  
  import (
  	"github.com/gin-gonic/gin"
  	"net/http"
  )
  
  // Login 定义接收数据的结构体
  type Login struct {
  	// binding:"required" 修饰的字段 若接收为空值 则报错 是必须字段
  	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
  	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
  }
  
  func main() {
  	r := gin.Default()
  	// JSON 绑定
  	r.POST("loginJson", func(context *gin.Context) {
  		// 声明接收的变量
  		var json Login
  		// 将request的body中的数据，自动按照JSON格式解析到结构体
  		if err := context.ShouldBindJSON(&json); err != nil {
  			// 返回错误信息
  			// gin.H封装了生成json数据的工具
  			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  			return
  		}
  		// 判断用户名密码是否正确
  		if json.User != "admin" || json.Password != "123456" {
  			context.JSON(http.StatusBadRequest, gin.H{"status": "304"})
  			return
  		}
  		context.JSON(http.StatusOK, gin.H{"status": "200"})
  	})
  	err := r.Run(":8080")
  	if err != nil {
  		return
  	}
  }
  ```

  ![image-20220417113234084](C:\Users\73554\AppData\Roaming\Typora\typora-user-images\image-20220417113234084.png)

#### 表单数据解析和绑定

```go
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    <form action="http://localhost:8080/loginForm" method="post" enctype="application/x-www-form-urlencoded">
        用户名<input type="text" name="username"><br>
        密码<input type="password" name="password">
        <input type="submit" value="提交">
    </form>
</body>
</html>
```

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Form 数据解析和绑定

type Login struct {
	// binding:"required" 修饰的字段 若接收为空值 则报错 是必须字段
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/loginForm", func(context *gin.Context) {
		// 声明接收的变量
		var form Login
		// Bind()默认解析并绑定form格式
		// 根据请求头中的content—type自动推断
		if err := context.Bind(&form); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if form.User != "admin" || form.Password != "123456" {
			context.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
```

![image-20220417174559693](C:\Users\73554\AppData\Roaming\Typora\typora-user-images\image-20220417174559693.png)

​			![image-20220417175311498](C:\Users\73554\AppData\Roaming\Typora\typora-user-images\image-20220417175311498.png)

#### URL数据解析和绑定

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Url 数据解析和绑定

type Login struct {
	// binding:"required" 修饰的字段 若接收为空值 则报错 是必须字段
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/:user/:password", func(context *gin.Context) {
		// 声明接收的变量
		var login Login
		// Bind()默认解析并绑定form格式
		// 根据请求头中的content—type自动推断
		if err := context.ShouldBindUri(&login); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if login.User != "admin" || login.Password != "123456" {
			context.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
```

![image-20220417184457127](C:\Users\73554\AppData\Roaming\Typora\typora-user-images\image-20220417184457127.png)

### Gin 渲染

#### 各种数据格式的响应

- json、结构体、XML、YAML类似于java的properties、ProtoBuf

  ```go
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
  ```

  ![image-20220417225132459](C:\Users\73554\AppData\Roaming\Typora\typora-user-images\image-20220417225132459.png)

![image-20220417225200060](C:\Users\73554\AppData\Roaming\Typora\typora-user-images\image-20220417225200060.png)

![image-20220417225229346](C:\Users\73554\AppData\Roaming\Typora\typora-user-images\image-20220417225229346.png)

![image-20220417225428674](C:\Users\73554\AppData\Roaming\Typora\typora-user-images\image-20220417225428674.png)

![image-20220417225607720](C:\Users\73554\AppData\Roaming\Typora\typora-user-images\image-20220417225607720.png)

#### HTML模板渲染

- gin支持加载HTML模板, 然后根据模板参数进行配置并返回相应的数据，本质上就是字符串替换

- LoadHTMLGlob()方法可以加载模板文件

  ```go
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
  ```

  ![image-20220417231038205](C:\Users\73554\AppData\Roaming\Typora\typora-user-images\image-20220417231038205.png)



#### 重定向





#### 同步异步











