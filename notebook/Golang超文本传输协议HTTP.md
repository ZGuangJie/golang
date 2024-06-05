## HTTP超文本传输协议：net/http

​	**超文本传输协议**（HTTP）是互联网上应用最为广泛的一种网络传输协议，是Web应用的应用层协议，定义==浏览器如何向Web服务器发送请求==，以及==Web服务器如何进行响应==。Golang内置得`net/http`包提供了HTTP**客户端**和**服务端**的实现。

`net/http`库主要由以下几个部分组成：

- **服务器端**：用于构建HTTP服务器。
- **处理器**：用于处理HTTP请求和响应。
- **客户端**：用于发起HTTP请求。
- **路由**：将请求路由映射到相应的处理器。

### 1、使用`net/http`构建HTTP服务器

​	构建服务器时需要做以下几件事情：

- **定义超时**
- **注册响应函数**：在响应函数里写响应的逻辑处理。
- **启动服务器监听端口**：接收客户端的请求。

```go
package main

import (
	"fmt"
	"net/http"
)
func main() {
	// 方式一：定义启动参数
	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
    
    // 设置路由
	http.HandleFunc("/test", httpResponse)
    
	// 开启监听
	fmt.Println("Server is listening on port 8080...")
    // 方式一启动
    if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed 	{
		fmt.Print("Fail to strat server...")
	}
	/*
	# 方式二：直接使用 http.ListenAndServe 启动
    if err := http.ListenAndServe(":8080", nil); err != nil {
	 	fmt.Print("Server run fail...")
	}
	*/ 
}
```



### 2、`net/http`处理HTTP请求和响应

​	使用http.HandleFunc()函数注册给定路由的处理函数。处理函数的参数，处理函数的签名通常是 `func(w http.ResponseWriter, r *http.Request)`。

​	处理函数中的参数`http.ResponseWriter` 和 `*http.Request` 是 Go 语言中处理 HTTP 请求和响应的关键接口。下面是关于这些包的信息：

#### 2.1 **`http.ResponseWriter`**

​	http.ResponseWriter是一个接口，用于构造 HTTP 响应。它包含以下方法：

- [**`Header()`**：用于设置或获取所有响应头信息。响应头包括内容类型、编码方式、缓存控制、Cookie 等](https://cloud.tencent.com/developer/article/1634278)[1](https://cloud.tencent.com/developer/article/1634278)。其作用是设置 HTTP 响应的内容类型（Content-Type），其类型包括：
    - `text/plain`：纯文本，不包含任何格式化或特殊样式。
    - `text/html`：HTML 格式的文本，用于构建网页。
    - `application/json`：JSON 格式的数据，用于传输结构化数据。
    - `image/jpeg`：JPEG 图像格式。
    - `application/pdf`：PDF 文档格式。
- **`Write([]byte) (int, error)`**：用于将数据写入响应实体。你可以通过这个方法将文本、JSON、HTML 等数据写入响应体。
- **`WriteHeader(statusCode int)`**：用于设置响应状态码。如果不调用该方法，默认状态码是 200 OK。

例如，你可以在处理器函数中使用 `http.ResponseWriter` 来构造响应，如下所示：

```go
func Home(w http.ResponseWriter, r *http.Request) {
    // 设置响应头
    w.Header().Set("Content-Type", "text/plain")
    // 写入响应实体：两种方式
    fmt.Fprintf(w, "Hello, World!")
    _, _ = w.Write([]byte("Welcome to my blog site"))
}
```
#### 2.2 **`*http.Request`**

​	*http.Request是一个结构体，封装了客户端发起的 HTTP 请求信息。它包含以下字段：

- **`Method`**：请求方法（GET、POST、PUT 等）。

- **`URL`**：请求的 URL。

- [**`Header`**：请求头信息，包括 Host、User-Agent、Accept、Accept-Encoding、Content-Length 等](https://www.runoob.com/http/http-messages.html)[2](https://www.runoob.com/http/http-messages.html)。

- **`Body`**：请求体，如果有的话，例如 POST 请求中的表单数据或 JSON 数据。

    获取Request中的信息：

    ```go
    // 获取GET访问方式中URL的参数信息: r *http.Request
    value := r.URL.Query().Get("params")
    url := r.URL.String() // 获取URL
    
    
    // r.ParseForm() 函数用于解析表单数据，生成的formData变量包含客户端提交的键值对，可以通过键访问各个表单值
    _ = r.ParseForm()
    formData := r.Form
    
    // 获取请求体（如果有）
    // 获取POST请求中的JSON数据，在Body中，使用r.Body获取，使用json库解析
    ```

### 3、`net/http`发起HTTP请求

​		可以直接使用`net/http`库发起HTTP请求，主要分为下面几个步骤：

- 确定URL

- 发送HTTP请求

- 延时关闭

- 设置超时参数，确保不会无限期阻塞

    ```go
    package main
    import (
        "fmt"
        "io"
        "net/http"
        "time"
    )
    func main() {
        client := &http.Client{
            Timeout: 10 * time.Second,
        }
    
        response, err := client.Get("http://localhost:8080/test")
        if err != nil {
            fmt.Println("Error", err)
            return
        }
        defer response.Body.Close()
        // 读取响应体
        body, err := io.ReadAll(response.Body)
        if err != nil {
            fmt.Println("Error reading response body:", err)
            return
        }
        fmt.Println("Response Body:", string(body))
    }
    ```

    

### 4、`net/http`请求路由映射到相应的处理器

​	在 `net/http` 库中，**路由**是指将收到的HTTP请求分派到相应的处理器函数，根据请求路径分配到不同的处理函数。路由的核心机制在于**请求路径**和**处理器**之间的映射。

​	在`net/http`库中，路由通过注册处理器函数或处理器对象来实现。以下是几种主要的路由方法：

- **`http.HandleFunc`**：用于注册一个处理器函数。
- **`http.Handle`**：用于注册一个实现了`http.Handler`接口的处理器对象。

```go
package main
import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    // 注册处理器对象
    http.Handle("/", http.FileServer(http.Dir("www")))
    // 注册处理器函数
    http.HandleFunc("/hello", helloHandler)
}
```

​	`http.Handle` 和 `http.HandleFunc` 都用于注册路由处理器，但它们之间有一些区别。让我们来看看这两者的不同之处：（**http.Handle就是先定义一个结构体，在结构体固定ServeHTTP接口里面实现处理逻辑，好处是可理解为面向对象编程**）

1. **`http.Handle`**：
    - `http.Handle` 接受两个参数：路由匹配的字符串和一个实现了 `Handler` 接口的结构体。
    - `Handler` 接口包含一个 `ServeHTTP(ResponseWriter, *Request)` 方法，用于处理接收到的请求。
    - [使用 `http.Handle` 时，您需要自定义一个结构体，实现 `Handler` 接口，并在该结构体的 `ServeHTTP`（必须要有） 方法中编写处理逻辑](https://blog.csdn.net/e891377/article/details/135713859)[1](https://blog.csdn.net/e891377/article/details/13571385)。
2. **`http.HandleFunc`**：
    - `http.HandleFunc` 接受两个参数：路由匹配的字符串和一个类型为 `func(ResponseWriter, *Request)` 的函数。
    - 这里的函数参数类型与 `Handler` 接口中的方法参数类型相同。
    - [使用 `http.HandleFunc` 时，您可以直接定义一个处理请求的函数，而无需创建一个结构体](https://blog.csdn.net/HYZX_9987/article/details/100017796)[2](https://blog.csdn.net/HYZX_9987/article/details/100017796)。

[总结：两者的功能相同，只是实现方式不同。`http.Handle` 需要定义一个实现了 `Handler` 接口的结构体，而 `http.HandleFunc` 则直接使用一个函数作为处理器。](https://www.cnblogs.com/leijiangtao/p/4509874.html)[3](https://www.cnblogs.com/leijiangtao/p/4509874.html)。