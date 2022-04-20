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

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 重定向

func main() {
	r := gin.Default()
	r.GET("/index", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "http:www.baidu.com")
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
```

#### 同步异步

- goroutine机制可以方便地实现异步处理

- 另外，在启动新的goroutine时，不应该使用原始上下文，必须使用它的只读副本

  ```go
  package main
  
  import (
  	"github.com/gin-gonic/gin"
  	"log"
  	"time"
  )
  
  func main() {
  	r := gin.Default()
  	// 1.异步
  	r.GET("async", func(context *gin.Context) {
  		// 需要一个副本
  		copyContext := context.Copy()
  		// 异步处理
  		go func() {
  			time.Sleep(3 * time.Second)
  			log.Println("异步执行：" + copyContext.Request.URL.Path)
  		}()
  		context.JSON(200, "同步~")
  	})
  	// 2.同步
  	r.GET("sync", func(context *gin.Context) {
  		time.Sleep(3 * time.Second)
  		log.Println("同步执行：" + context.Request.URL.Path)
  		context.JSON(200, "异步")
  	})
  	err := r.Run(":8080")
  	if err != nil {
  		return
  	}
  }
  ```

### Gin 中间件

#### 全局中间件

- 所有请求都经过此中间件

  ```go
  package main
  
  import (
  	"fmt"
  	"github.com/gin-gonic/gin"
  	"time"
  )
  
  // 全局中间件
  
  // MiddleWare 定义中间件
  func MiddleWare() gin.HandlerFunc {
  	return func(context *gin.Context) {
  		t := time.Now()
  		fmt.Println("中间件开始执行...")
  		// 设置变量到 context 的 key 中，可以通过 get()获取
  		context.Set("request", "中间件")
  		status := context.Writer.Status()
  		fmt.Println("中间件执行结束!", status)
  		t2 := time.Since(t)
  		fmt.Println("time:", t2)
  	}
  }
  
  func main() {
  	r := gin.Default()
  	// 注册中间件
  	r.Use(MiddleWare())
  	// {} 代码规范
  	{
  		r.GET("/middleware", func(context *gin.Context) {
  			// 取值
  			request, _ := context.Get("request")
  			fmt.Println("request:", request)
  			// 页面接收
  			context.JSON(200, gin.H{"request": request})
  		})
  	}
  	err := r.Run(":8080")
  	if err != nil {
  		return
  	}
  }
  ```

#### Next()方法

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// MiddleWare 定义中间件
func MiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行...")
		// 设置变量到 context 的 key 中，可以通过 get()获取
		context.Set("request", "中间件")
		// 执行函数
		context.Next()
		// 中间件执行完的后续动作
		status := context.Writer.Status()
		fmt.Println("中间件执行结束!", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func main() {
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 注册中间件
	r.Use(MiddleWare())
	// {} 代码规范
	{
		r.GET("/middleware", func(context *gin.Context) {
			// 取值
			request, _ := context.Get("request")
			fmt.Println("request:", request)
			// 页面接收
			context.JSON(200, gin.H{"request": request})
		})
	}
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
```

#### 局部中间件

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 局部中间件

// MiddleWare 定义中间件
func MiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行...")
		// 设置变量到 context 的 key 中，可以通过 get()获取
		context.Set("request", "中间件")
		// 执行函数
		context.Next()
		// 中间件执行完的后续动作
		status := context.Writer.Status()
		fmt.Println("中间件执行结束!", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func main() {
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 局部中间件
	r.GET("/middleware", MiddleWare(), func(context *gin.Context) {
		// 取值
		request, _ := context.Get("request")
		fmt.Println("request:", request)
		// 页面接收
		context.JSON(200, gin.H{"request": request})
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
```

#### 中间件练习

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 中间件
func myTime(ctx *gin.Context) {
	start := time.Now()
	ctx.Next()
	// 统计时间
	since := time.Since(start)
	fmt.Println("累计用时：", since)
}

func shopIndexHandler(ctx *gin.Context) {
	time.Sleep(5 * time.Second)
	ctx.String(http.StatusOK, "shopIndexHandler~~~")
}

func shopHomeHandler(ctx *gin.Context) {
	time.Sleep(3 * time.Second)
	ctx.String(http.StatusOK, "shopHomeHandler~~~")
}

func main() {
	r := gin.Default()
	// 注册中间件
	r.Use(myTime)
	shoppingGroup := r.Group("shopping")
	{
		shoppingGroup.GET("/index", shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
```

### 会话控制

#### Cookie 介绍

- HTTP是无状态协议，服务器不能记录浏览器的访问状态，也就是说服务器不能区分两次请求是否由同一个客户端发出
- Cookie就是解决HTTP协议无状态的方案之一，中文是小甜饼的意思
- Cookie实际上就是服务器保存在浏览器上的一段信息。浏览器有了Cookie之后，每次向服务器发送请求时都会同时将该信息发送给服务器，服务器收到请求后，就可以根据该信息处理请求
- Cookie由服务器创建，并发送给浏览器，最终由浏览器保存

**用途**

- 测试服务端发送cookie给客户端，客户端请求时携带cookie

#### Cookie 使用

- 测试服务端发送cookie给客户端，客户端请求时携带cookie

  ```go
  package main
  
  import (
  	"fmt"
  	"github.com/gin-gonic/gin"
  	"net/http"
  )
  
  func main() {
  	r := gin.Default()
  	// 服务器要给客户端cookie
  	r.GET("cookie", func(ctx *gin.Context) {
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
  ```

#### Cookie 练习

- 模拟实现权限验证中间件

  - 有2个路由，login和home
  - login用于设置cookie
  - home是访问查看信息的请求
  - 在请求home之前，先跑中间件代码，检验是否存在cookie

- 访问home，会显示错误，因为权限校验未通过

  ```go
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
  		ctx.JSON(http.StatusUnauthorized, gin.H{"mes": "cookie 鉴权失败"})
  		// 若验证不通过，不再调用后续的函数处理
  		ctx.Abort()
  		return
  	}
  }
  
  func main() {
  	r := gin.Default()
  	r.GET("/login", func(ctx *gin.Context) {
  		// 设置 cookie(domain 如果设置了 localhost 浏览器不能通过 127.0.0.1 来访问 否则 cookie 会设置失败)
  		ctx.SetCookie("abc", "123", 60, "/", "localhost", false, true)
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
  ```

  ![image-20220419141324729](https://tva1.sinaimg.cn/large/e6c9d24ely1h1f05a9ljzj213s088wgf.jpg)

关于 setCookie 时 domain 设置的是 localhost，启动之后通过 127.0.0.1/login 访问来设置 cookie 失败问题的原因我从网上找了一些信息，可以参考一下：

```go
和代码无关，我访问地址有问题，我访问的是 127.0.0.1:8001，应该是localhost:8000 127.0.0.1通常是分配给“环回”或本地接口的IP地址。这是一个只能在同一主机内通信的“假”网络适配器。当您希望具有网络功能的应用程序仅为同一主机上的客户机提供服务时，通常会使用这种方法。在127.0.0.1上监听连接的进程将只接收该套接字上的本地连接。

“localhost”通常是127.0.0.1 IP地址的主机名。它通常在/etc/hosts中设置(或者在%WINDIR%下的等效窗口名为“hosts”)。您可以像使用任何其他主机名一样使用它—尝试“ping localhost”，看看它是如何解析为127.0.0.1的。

0.0.0.0有几个不同的含义，但是在本文中，当服务器被告知监听0.0.0.0时，这意味着“监听每个可用的网络接口”。从服务器进程的角度来看，IP地址为127.0.0.1的环回适配器与机器上的任何其他网络适配器一样，因此被告知监听0.0.0.0的服务器也将接受该接口上的连接。
```

#### 缺点

- 不安全，明文
- 增加带宽消耗
- 可以被禁用
- cookie有上限

#### Sessions

gorilla/sessions为自定义session后端提供cookie和文件系统session以及基础结构。

主要功能是：

- 简单的API：将其用作设置签名（以及可选的加密）cookie的简便方法。
- 内置的后端可将session存储在cookie或文件系统中。
- Flash消息：一直持续读取的session值。
- 切换session持久性（又称“记住我”）和设置其他属性的便捷方法。
- 旋转身份验证和加密密钥的机制。
- 每个请求有多个session，即使使用不同的后端也是如此。
- 自定义session后端的接口和基础结构：可以使用通用API检索并批量保存来自不同商店的session。

代码：

```go
package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

// sessions

// 初始化一个cookie存储对象
var store = sessions.NewCookieStore([]byte("test-secret"))

func main() {
	http.HandleFunc("/save", SaveSession)
	http.HandleFunc("/get", GetSession)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP server failed, err:", err)
		return
	}
}

func SaveSession(w http.ResponseWriter, r *http.Request) {
	// 获取一个session对象 session-name是session的名字
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 在session中存储值
	session.Values["foo"] = "bar"
	session.Values[42] = 43
	// 保存更改
	_ = session.Save(r, w)
}

func GetSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	foo := session.Values["foo"]
	fmt.Println(foo)
}
```

### 参数验证

#### 结构体验证

用gin框架的数据验证，可以不用解析数据，减少if else，会简洁许多。

```go
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
```

#### 自定义验证

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// Login 对绑定解析到结构体上的参数，自定义验证功能
// 比如我们需要对URL的接收参数进行判断，判断用户名是否为root如果是root通过否则返回false
type Login struct {
	User     string `uri:"user" validate:"checkName"`
	Password string `uri:"password"`
}

// 自定义验证函数
func checkName(fl validator.FieldLevel) bool {
	if fl.Field().String() != "root" {
		return false
	}
	return true
}

func main() {
	r := gin.Default()
	validate := validator.New()
	r.GET("/:user/:password", func(ctx *gin.Context) {
		var login Login
		// 注册自定义函数，与struct tag 关联起来
		err := validate.RegisterValidation("checkName", checkName)
		if err := ctx.ShouldBindUri(&login); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = validate.Struct(login)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println(err)
			}
			return
		}
		fmt.Println("success")
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
```

#### 自定义验证 V10

Validator 是基于 tag（标记）实现结构体和单个字段的值验证库，它包含以下功能：

- 使用验证 tag（标记）或自定义验证器进行跨字段和跨结构体验证。
- 关于 slice、数组和 map，允许验证多维字段的任何或所有级别。
- 能够深入 map 键和值进行验证。
- 通过在验证之前确定接口的基础类型来处理类型接口。
- 处理自定义字段类型（如 sql 驱动程序 Valuer）。
- 别名验证标记，它允许将多个验证映射到单个标记，以便更轻松地定义结构体上的验证。
- 提取自定义的字段名称，例如，可以指定在验证时提取 JSON 名称，并在生成的 FieldError 中使用该名称。
- 可自定义 i18n 错误消息。
- Web 框架 gin 的默认验证器。

**安装**

> ```go
> go get github.com/go-playground/validator/v10
> ```

**变量验证**

Var 方法使用 tag（标记）验证方式验证单个变量。

```go
func (*validator.Validate).Var(field interface{}, tag string) error
```

它接收一个 interface{} 空接口类型的 field 和一个 string 类型的 tag，返回传递的非法值得无效验证错误，否则将 nil 或 ValidationErrors 作为错误。如果错误不是 nil，则需要断言错误去访问错误数组，例如：

```go
validationErrors := err.(validator.ValidationErrors)
```

如果是验证数组、slice 和 map，可能会包含多个错误。

示例代码：

```go
func main() {
  validate := validator.New()
  // 验证变量
  email := "admin#admin.com"
  err := validate.Var(email, "required,email")
  if err != nil {
    validationErrors := err.(validator.ValidationErrors)
    fmt.Println(validationErrors)
    // output: Key: '' Error:Field validation for '' failed on the 'email' tag
    return
  }
}
```

**结构体验证**

结构体验证结构体公开的字段，并自动验证嵌套结构体，除非另有说明。

```go
func (*validator.Validate).Struct(s interface{}) error
```

它接收一个 interface{} 空接口类型的 s，返回传递的非法值得无效验证错误，否则将 nil 或 ValidationErrors 作为错误。如果错误不是 nil，则需要断言错误去访问错误数组，例如：

```go
validationErrors := err.(validator.ValidationErrors)
```

实际上，Struct 方法是调用的 StructCtx 方法，因为本文不是源码讲解，所以此处不展开赘述，如有兴趣，可以查看源码。

示例代码：

```go

```







#### 多语言翻译验证













