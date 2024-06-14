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

​		在golang里面的模板与前端的html、css、js遵循的规则一致，`不一样的是`有一些基础的语法，相当于占位符，从而在后端完成变量的替换。

#### 2.2 解析模板文件

​		定义好模板文件以后，使用下面的三个方法解析模板，可以得到模板对象：

##### 2.2.2 .ParseFiles

```go
package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
)
func pading(w http.ResponseWriter, r *http.Request) {
    // 1.定义模板 Details see dir templet
    // 2.解析模板
    t, err := template.ParseFiles("./templet/base.html")
    if err != nil {
        log.Fatal("文件解析错误: ", err)
    }
    // 3.渲染模板
    t.Execute(w, "Hello")
}
func main() {
    http.HandleFunc("/base", pading)

    err := http.ListenAndServe(":9000", nil)
    fmt.Println("Server Start.")
    if err != nil {
        log.Fatal("Fail to start service: ", err)
    }
}
/*
	将"Hello"字符串写入 w http.ResponseWriter，传到客户端。
*/
```

​		可以使用`func New(name string) *Template`函数创建一个名为`name`的模板，然后对其调用上面的方法去解析模板字符串或模板文件。

```go
// 1.定义模板 Details see dir templet
htmlByte, err := os.ReadFile("./templet/ageAdd.html")
if err != nil {
    fmt.Println("read html failed, err:", err)
    return
}
// 2.解析模板
t, err := template.New("AgeAdd").Funcs(template.FuncMap{"AgeAdd": AddAge}).Parse(string(htmlByte))
if err != nil {
    log.Fatal("解析模板错误: ", err)
    }
```



#### 2.3 渲染模板文件

​		渲染模板简单来说就是**使用数据去填充模板**，当然实际上可能会复杂很多。

##### 2.3.1 .Excute

​		`Execute`用于执行一个模板（通常是通过`New`或`Parse`创建的模板）。

```go
func (t *Template) Execute(wr io.Writer, data interface{}) error
```

##### 2.3.1 .ExcuteTemplate

​		`ExecuteTemplate`用于执行一个命名模板。当你有多个命名模板（通过`{{define "name"}}...{{end}}`定义）并且需要选择其中一个进行执行时使用。

```go
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
```

### 3、模板语法

#### 3.1 {{.}}

​		`{{.}}`中的`.`表示传入的对象。和在golang中一样使用。

```go
// In Golang 首字母小写在外访问不到
type stu struct {
	Name string
	Age  int16
	Sex  bool
}

//In templet
<p>{{ .Name }}</p>
<p>{{ .Age }}</p>
<p>{{ .Sex }}</p>
```

#### 3.2 变量定义 与 注释

​		在templet模板中的变量，语句均在`{{}}`中操作。

##### 3.2.1 变量

```html
// 使用$来命名变量，全在{{}}操作.
{{$intersting := .Age}}
<p>{{$intersting}}</p>
```

##### 3.2.2 注释

`{{/* a comment */}}`执行时会忽略。可以多行。注释不能嵌套，并且必须紧贴分界符始止。

#### 3.3 条件判断

​		Go模板中的条件判断有以下三种:

```go
{{if pipeline}} T1 {{end}}
{{if pipeline}} T1 {{else}} T0 {{end}}
{{if pipeline}} T1 {{else if pipeline}} T0 {{end}}
```

```html
<body>
    <hr>
    {{ $flag := true }}
    {{if $flag}}
    {{ with .slice }}
    {{ range $id, $val := . }}
    {{ $id }} 
    {{ $val }}
    {{ end }}
    {{ end }}
    {{ end }}
</body>
```



#### 3.4 range 遍历

​	使用range对golang里面的 可迭代结构 进行遍历:

```html
<body>
    {{ with .stu }}
    <p>{{ .Name }}</p>
    <p>{{ .Age }}</p>
    <p>{{ .Sex }}</p>
    {{end}}
    <hr>
    {{ with .slice }}
    {{ range $id, $val := . }}
    {{ $id }} 
    {{$val}}
    {{ end }}
    {{ end }}
</body>
```

#### 3.5 基本预定义函数

​		在templet模板中提供一些简单的函数，and与、or或函数，index索引取值函数，**printf输出函数**等。

##### 3.5.1 预定义全局函数：

```template
and
    函数返回它的第一个empty参数或者最后一个参数；
    就是说"and x y"等价于"if x then y else x"；所有参数都会执行；
or
    返回第一个非empty参数或者最后一个参数；
    亦即"or x y"等价于"if x then x else y"；所有参数都会执行；
not
    返回它的单个参数的布尔值的否定
len
    返回它的参数的整数类型长度
index
    执行结果为第一个参数以剩下的参数为索引/键指向的值；
    如"index x 1 2 3"返回x[1][2][3]的值；每个被索引的主体必须是数组、切片或者字典。
print
    即fmt.Sprint
printf
    即fmt.Sprintf
println
    即fmt.Sprintln
html
    返回与其参数的文本表示形式等效的转义HTML。
    这个函数在html/template中不可用。
urlquery
    以适合嵌入到网址查询中的形式返回其参数的文本表示的转义值。
    这个函数在html/template中不可用。
js
    返回与其参数的文本表示形式等效的转义JavaScript。
call
    执行结果是调用第一个参数的返回值，该参数必须是函数类型，其余参数作为调用该函数的参数；
    如"call .X.Y 1 2"等价于go语言里的dot.X.Y(1, 2)；
    其中Y是函数类型的字段或者字典的值，或者其他类似情况；
    call的第一个参数的执行结果必须是函数类型的值（和预定义函数如print明显不同）；
    该函数类型值必须有1到2个返回值，如果有2个则后一个必须是error接口类型；
    如果有2个返回值的方法返回的error非nil，模板执行会中断并返回给调用模板执行者该错误；
```

##### 3.5.2 比较函数

布尔函数会将任何类型的零值视为假，其余视为真。

```go
eq      如果arg1 == arg2则返回真
ne      如果arg1 != arg2则返回真
lt      如果arg1 < arg2则返回真
le      如果arg1 <= arg2则返回真
gt      如果arg1 > arg2则返回真
ge      如果arg1 >= arg2则返回真
```

​		`eq（只有eq）`可以接受2个或更多个参数

#### 3.6 自定义函数

​		Go的模板可以支持自定义函数。分两步

1. 定义函数；
2. 使用Funcs绑定函数。

```go
package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "os"
)

type stu struct {
    Name string
    Age  int16
    Sex  bool
}

func AddAge(age interface{}) int16 {
    var add int16
    if age, ok := age.(int16); ok {
        add = age + 1
    }
    return add
}
func ageAdd(w http.ResponseWriter, r *http.Request) {
    // 1.定义模板 Details see dir templet
    htmlByte, err := os.ReadFile("./templet/ageAdd.html")
    if err != nil {
        fmt.Println("read html failed, err:", err)
        return
    }
    // 2.解析模板
    t, err := template.New("AgeAdd").Funcs(template.FuncMap{"AgeAdd": AddAge}).Parse(string(htmlByte))
    if err != nil {
        log.Fatal("解析模板错误: ", err)
    }
    // 渲染模板
    student := stu{
        Name: "zhu",
        Age:  26,
        Sex:  true,
    }
    t.Execute(w, student)
}
func main() {
    http.HandleFunc("/base", pading)
    http.HandleFunc("/age", ageAdd)

    err := http.ListenAndServe(":9000", nil)
    fmt.Println("Server Start.")
    if err != nil {
        log.Fatal("Fail to start service: ", err)
    }
}
```

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

