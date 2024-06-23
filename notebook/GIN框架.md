### GIN框架

​		web浏览器所有的交互都是一次**请求(request)**，一次**响应(response)**。

![](https://cdn.jsdelivr.net/gh/ZGuangJie/GoPicture/golang/202406141449428.png)

#### 1、gin框架安装与使用

​		安装gin框架：

```go
go get -u github.com/gin-gonic/gin
```

##### RESTful API:

​		REST与技术无关，代表的是一种软件架构风格，REST是Representational State Transfer的简称，中文翻译为“表征状态转移”或“表现层状态转化”。

推荐阅读[阮一峰 理解RESTful架构](http://www.ruanyifeng.com/blog/2011/09/restful.html)

​		简单来说，REST的含义就是客户端与Web服务器之间进行交互的时候，使用HTTP协议中的4个请求方法代表不同的动作。

- `GET`用来获取资源
- `POST`用来新建资源
- `PUT`用来更新资源
- `DELETE`用来删除资源。

​		URI（统一资源定位符）不应该有动词，因为“资源”是一个实体，用`HTTP不同的请求方法`来表示动作。



#### 2、Gin框架下的渲染

​		与使用net/http框架一样也需要创建模板、解析模板和渲染模板。

##### 2.1 创建模板

​		一般是前端工程师做的事情。

##### 2.2 解析模板

​		通常使用两个方法 .LoadHTMLFiles() 或 .LoadHTMLGlob()，返回解析后文件的句柄。前者是通过**文件名字**解析文件，后者是通过**正则匹配**解析文件。

##### 2.3 渲染文件

​		和net/http库类似，建立路由处理函数处理对应的路由请求。不同的是，可以区分HTTP不同的请求方式，如Get、Post、Put等等。

```go
r.GET("/index", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{
        "title": "",
    })
})
/*
	请求的参数全都去 c 里去找。
*/
```

##### 2.4 渲染静态文件

​		使用下面三个函数来解析静态文件。

```go
r.Static("/assets", "./assets")
r.StaticFS("/more_static", http.Dir("my_file_system"))
r.StaticFile("/favicon.ico", "./resources/favicon.ico")
```

**`router.Static`**：将一个本地文件夹作为静态资源目录，提供目录下的所有文件。

**`router.StaticFS`**：与 `Static` 类似，但可以自定义文件系统（例如嵌入文件系统或自定义文件系统）。

**`router.StaticFile`**：提供单个静态文件。

```go
/*
文件目录树
>static
  >css
    >css2.css
    >style.css
    >bootstrap.min.css
    >font-awesome.min.css
    >themify-icons.css
  >font
  >image
*/
// 解析static目录下所有的文件
r.Static("/static", "./static")
```

​		在html文件里使用方式如下，以static为根目录访问

```html
<!-- Latest Bootstrap min CSS -->
<link rel="stylesheet" href="/static/css/bootstrap.min.css">		
<!-- Google Font -->
<!-- Font Awesome CSS -->
<link rel="stylesheet" href="/static/css/font-awesome.min.css">
<link rel="stylesheet" href="/static/css/themify-icons.css">
```

#### 3、获取请求中的参数

​		不同的请求方式获取参数的方式稍有不同，主要是Get请求方式、Post请求方式。

##### 3.1 Get请求方式

​		获取Query String里面的参数，相当于net/http下的 router.HandleFunc(“/index”, func())。

```go
r.GET("/index", func(c *gin.Context) {
    // 方式一
    // name, ok := c.GetQuery("name")
    // if !ok {
    // 	name = "Cannot find name"
    // }
    
    // 方式二
    name := c.DefaultQuery("name", "Cannot find name")
    // fmt.Printf("%v", name)
    c.JSON(http.StatusOK, gin.H{"message": name})
})
```

##### 3.2 Post请求方式

​		Post一般是接收前端页面提交的表单。

```go
r.POST("/index", func(c *gin.Context) {
    // 方式一
    // name := c.PostForm("name")
    // pwd := c.PostForm("password")
    
    // 方式二
    name := c.DefaultPostForm("nae", "None")
    pwd := c.DefaultPostForm("pasword", "XXX")
    c.JSON(http.StatusOK, gin.H{
        "message": pwd,
        "name":    name,
    })
})
```

##### 3.3 参数绑定

​		.ShouldBind()方法，将form表单里的参数对应存到一个结构体里。注意**结构体 tag 是form**。

```go
type students struct {
    Name string `form:"name"`
    Age  int16  `form:"age"`
}
r.POST("/index", func(c *gin.Context) {
    var stu students
    err := c.ShouldBind(&stu)
    if err != nil {
        c.JSON(http.StatusBadGateway, "Params delivery ERROR!")
    } else {
        fmt.Printf("%#v\n", stu)
        c.JSON(http.StatusOK, stu)
    }
})
```

​		html表单里的参数要对应：

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
    </head>
    <body>
        <form action="/index" method="post">
            <div>
                <p>name:</p>
                <<!-- name 要对应到后端的结构体中 -->
                <input type="text" name="name" required>
            </div>
            <div>
                <p>age:</p>
                <input type="text" name="age" required>
            </div>

            <button>提交</button>
        </form>
    </body>
</html>
```

##### 3.4 获取URI

​		URI是**统一资源定位符**，其实就是**域名后面**的那些个字符串。

```go
r.GET("/index", func(c *gin.Context) {
    // fmt.Printf("%v", name)
    c.JSON(http.StatusOK, c.Request.URL)
})
/*
下面这些参数都可以获取到：
{"Scheme":"","Opaque":"","User":null,"Host":"","Path":"/index","RawPath":"","OmitHost":false,"ForceQuery":false,"RawQuery":"name=hello","Fragment":"","RawFragment":""}
*/
```

##### 3.4 文件上传

​		在前端上传文件，发送到后端存储。

```go
r.GET("/upload", func(c *gin.Context) {
    // fmt.Printf("%v", name)
    c.HTML(http.StatusOK, "upload.html", nil)
})
r.MaxMultipartMemory = 8 << 20 // 8 MiB
r.POST("/upload", func(c *gin.Context) {
    // 单文件
    file, _ := c.FormFile("file")
    log.Println(file.Filename)

    dst := filepath.Join("./uploadFiles", file.Filename)
    // 上传文件至指定的完整文件路径
    c.SaveUploadedFile(file, dst)

    c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
})
```

#### 4、重定向

​		常用在支付完成后，跳转到原页面，出现404后跳转等

##### 4.1 外部重定向

​		跳转到外部网页。

```go
r.GET("/redirect", func(c *gin.Context) {
    c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
})
```

##### 4.2 内部重定向

​	修改请求的URI，继续后续的处理。

```go
r.GET("/test", func(c *gin.Context) {
    // 在访问 /test 时直接 转到 /upload
    c.Redirect(http.StatusFound, "/upload")
})
```



#### 5、路由和路由组

​		客户端每次访问的路径，每条路径对应着一个处理函数，也可以说，每个路由对应一个处理函数。

##### 5.1 路由

​		包括常用四个请求方式，.GET()、.POST()、.PUT()、.DELETE()。

```go
// 此 handler 将匹配 /user/X，即/user/下一层的所有路由（只有一层）， 但不会匹配 /user/ 或者 /user
router.GET("/user/:name", func(c *gin.Context) {
    name := c.Param("name")
    c.String(http.StatusOK, "Hello %s", name)
})

// 此 handler 将匹配 /user/john/ 和 /user/john/send （可以访问到/user/下的所有路由）
// 如果没有其他路由匹配 /user/john，它将重定向到 /user/john/
router.GET("/user/:name/*action", func(c *gin.Context) {
    name := c.Param("name")
    action := c.Param("action")
    message := name + " is " + action
    c.String(http.StatusOK, message)
})
```



##### 5.2 路由组

​		我们可以将拥有**共同URL前缀的路由**划分为一个路由组。习惯用一对包裹同组的路由，这只是为了看着清晰，功能上没区别。使用r.Group()，通常我们将**路由分组**用在划分业务逻辑或划分API版本。

```go
// 简单的路由组: v1
v1 := router.Group("/v1")
{
    // 相当于访问 /v1/login, ect
    v1.POST("/login", loginEndpoint)
    v1.POST("/submit", submitEndpoint)
    v1.POST("/read", readEndpoint)
}

// 简单的路由组: v2
v2 := router.Group("/v2")
{
    v2.POST("/login", loginEndpoint)
    v2.POST("/submit", submitEndpoint)
    v2.POST("/read", readEndpoint)
}

```

#### 6、Gin中间件

​		Gin框架允许开发者在处理请求的过程中，加入自己的钩子函数（Hook）函数。这个钩子函数就叫中间件，中间件适合**处理一些公共的业务逻辑**，比如登录认证、权限认证、数据分页、记录日志、耗时统计等。

![中间件](https://cdn.jsdelivr.net/gh/ZGuangJie/GoPicture/golang/202406191100897.png)

##### 6.1 中间件函数定义

​		类似于处理函数前在加一个处理函数，定义方式一样，可以使用 .Use()全局定义。

C.Next()执行后面的处理函数，就是函数调用，压栈的过程。

```go
r := gin.New()
// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
r.Use(gin.Logger())
// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
r.Use(gin.Recovery())
r.Static("/static", "./static")
// 解析HTML模板
r.LoadHTMLGlob("./template/*.html")

// 访问/login 时，先经过 verify 中间件
r.GET("/login", verify(c *gin.Context), func(c *gin.Context) {
    c.HTML(http.StatusOK, "login.html", nil)
})
```

##### 6.2 闭包使用中间件

​		闭包中间件的使用场合

1. **参数化中间件**: 当你需要为中间件提供不同的配置参数时，闭包中间件非常有用。例如，可以创建日志记录中间件，并根据不同的路由或应用部分使用不同的日志前缀或日志级别。
2. **共享状态**: 闭包中间件可以保持状态，这对于需要在多个请求中共享状态或数据的场合非常有用。例如，统计访问计数或限制请求速率的中间件。
3. **动态配置**: 通过闭包，你可以动态地创建和配置中间件，而无需在全局作用域中定义所有中间件。

##### 6.3 为某个路由组注册中间件

​		使该路由组均经过固定的中间件。

```go
login := r.Group("/index")
// loginVerify(c *gin.Context),自定义中间件处理函数
login.Use(loginVerify)
{
    login.POST("/submit", func(c *gin.Context) {
        message, ok := c.Get("example")
        if !ok {
            message = "Null"
        }
        c.JSON(http.StatusOK, gin.H{"URI": message})
    })
    login.POST("/error", func(c *gin.Context) {
        c.String(http.StatusOK, "用户名或者密码错误")
    })
}
```



##### 6.4 跨中间件获取值

​		在中间件中使用c.Set(“”, “”)设置一个值，可以使用 c.Get() 在后面的路由处理函数中拿到。

```go
c.Set("example", "Recived")
// --------------处理函数---------------- //
login.POST("/submit", func(c *gin.Context) {
    message, ok := c.Get("example")
    if !ok {
        message = "Null"
    }
    c.JSON(http.StatusOK, gin.H{"URI": message})
})
```



##### 6.5 中间件并发安全

​		为了并发安全，在中间件内新开一个gorotine处理请求时，应该传入 gin.Context的副本，否则在后续的处理函数中拿到的 c *gin.Context是不安全的。

​		

#### 7、Object Relational Mapping（ORM，对象关系映射）



GORM Model定义

​		GORM 通过将 Go 结构体（Go structs） 映射到数据库表来简化数据库交互。 了解如何在GORM中定义模型，是充分利用GORM全部功能的基础。





#### 8、项目构成







#### http状态码：