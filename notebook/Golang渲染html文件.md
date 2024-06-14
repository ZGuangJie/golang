## Golang之html文件渲染

### 1、模板与渲染

​		在一些**前后端不分离**的Web架构中，我们通常需要**在后端将一些数据渲染到HTML文档中**，从而实现动态的网页效果。(布局和样式大致一样，展示的内容不一样)

- 模板可以理解为**事先定义好的HTML文档文件**，
- 渲染的作用机制可以简单理解为**文本替换操作–**使用相应的数据去替换HTML文档中**事先定义好的标记**。

#### 1.1 Go语言模板引擎

​		Go语言内置了文本模板引擎`text/template`和用于HTML文档的`html/template`。它们的作用机制可以简单归纳如下：

1. 模板文件通常定义为`.tmpl`和`.tpl`为后缀（也可以使用其他的后缀），必须使用`UTF8`编码。
2. 模板文件中使用`{{`和`}}`包裹和标识需要传入的数据。
3. 传给模板这样的数据就可以通过点号（`.`）来访问，如果数据是复杂类型的数据，可以通过{ { .FieldName }}来访问它的字段。
4. 除`{{`和`}}`包裹的内容外，其他内容均不做修改原样输出。

​		只替换 {{ }} 里面的内容。

### 2、模板引擎的使用

#### 2.1 定义模板文件

#### 2.2 解析模板文件

#### 2.3 渲染模板文件



### 3、模板语法

#### 3.1 {{.}}

#### 3.2 变量定义 与 注释

#### 3.3 条件判断

#### 3.4 range 遍历

#### 3.5 基本预定义函数

#### 3.6 自定义函数



### 4、模板的嵌套

​		我们可以在template中嵌套其他的template。这个template可以是单独的文件，也可以是通过`define`定义的template。

#### 4.1 两种情况

1. 单独的template文件嵌套；
2. 通过define自定义的template嵌套；

```go
```

#### 4.2 block 模板的继承

​		使用`block`定义一组根模板，然后在子模版里对**块模板重新定义**。在基础模板里 **{{block "content" . }}{{end}}** 使用预留一个代码块，在子模版里定义 填充的内容，需要指定继承的哪个基础模板**{{templates/base.tmpl}}**。

```go
```







​		

#### 模板继承

#### 2、使用自带库 html/template

​	在Go语言中，可以使用`html/template`包加载并渲染一个HTML模板。具体步骤如下：

1. ==创建html模板==

    ​		模板文件由文本和控制结构组成，最简单的空值结构是动作（action），由一对大括号包围，如{{.FileName}}，输出字段值。

    ```html
    <!DOCTYPE html>
    <html lang="en">
        <head>
            <meta charset="UTF-8">
            <title>{{ .Title }}</title>
        </head>
        <body>
            <h1>{{ .Heading }}</h1>
            <p>{{ .Content }}</p>
            <ul>
                {{ range .Items }}
                <li>{{ . }}</li>
                {{ end }}
            </ul>
        </body>
    </html>
    ```

    

2. ==解析html模板==

    ​		通过`template.ParseFiles()`或`template.ParseGlob()`解析模板文件。

    ```go
    func renderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
        t, err := template.ParseFiles(tmpl)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        err = t.Execute(w, data)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
    ```

    

3. ==将数据渲染到模板==

    ​	调用`Execute`方法将数据填充到模板中。

    ```go
    func renderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
    	t, err := template.ParseFiles(tmpl)
    	if err != nil {
    		http.Error(w, err.Error(), http.StatusInternalServerError)
    		return
    	}
    	err = t.Execute(w, data)
    	if err != nil {
    		http.Error(w, err.Error(), http.StatusInternalServerError)
    	}
    }
    ```

    

4. ==开启net/http服务==

    ```go
    func main() {
        // 方式一：定义启动参数
        server := &http.Server{
            Addr:    ":8080",
            Handler: nil,
        }
    
        // 设置路由
        http.HandleFunc("/test", renderTemplate)
    
        // 开启监听
        fmt.Println("Server is listening on port 8080...")
        // 方式一启动
        if err := server.ListenAndServe(); err != nil && err != 					http.ErrServerClosed{
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

    

5. 

