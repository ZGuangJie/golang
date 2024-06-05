## Golang之html文件渲染

### 1、使用自带库 html/template

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

