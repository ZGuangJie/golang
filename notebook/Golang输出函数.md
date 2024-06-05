

## Golang - - 格式化和输出函数

​	在 Go 语言中，`fmt` 包提供了一组用于格式化和打印的函数。这些函数可以分为两类：==生成字符串==和==直接输出到标准输出==（如控制台）。

### 1、生成字符串的函数

​	下面这些函数生成格式化后的字符串，但不会直接输出到标准输出。

#### 1.1 fmt.Sprint 

​	`fmt.Sprint` 连接所有传入的参数，并生成一个字符串。

```go
s := fmt.Sprint("Hello", " ", "world!")
fmt.Println(s) // 输出: Hello world!
```

#### 1.2 fmt.Sprintf

​	`fmt.Sprintf` 根据格式说明符生成格式化后的字符串。

```go
age := 26
name := "zhu"
s := fmt.Sprintf("Age:%d, Name:%s", age, name)
fmt.Println(s) // 输出: Age:26, Name:zhu
```

#### 1.3 fmt.Sprintln 

​	`fmt.Sprintln` 连接所有传入的参数，并在每个参数之间添加空格，最后添加一个换行符。

```go
s := fmt.Sprintln("Hello", "world!")
fmt.Println(s) // 输出: Hello world!
```

### 2、直接输出到标准输出

​	下面这些函数直接将格式化后的字符串输出到标准输出。

#### 2.1 fmt.Print

`fmt.Print` 连接所有传入的参数，并直接输出到标准输出。

```go
fmt.Print("Hello", " ", "World!")
//output: Hello World!
```

#### 2.2 fmt.Printf

`fmt.Printf` 根据格式说明符，将**格式化后的字符串**直接输出到标准输出。

```go
age := 26
name := "zhu"
fmt.Sprintf("Age:%d, Name:%s", age, name)
//output: Age:26, Name:zhu
```

#### 2.3 fmt.Print

`fmt.Println` 连接所有传入的参数，**参数之间加空格，末尾加换行符**，然后接输出到标准输出。

```go
fmt.Print("Hello", " ", "World!")
//output: Hello World!
```

### 3、输输出到非标准输出（文件）

​	除了可以使用 `fmt` 包输出到标准输出（控制台）之外，还可以将格式化后的内容输出到文件。要实现这一点，需要使用 `os` 包打开或创建文件，然后使用 `fmt` 包的相关函数将内容写入文件。

```go
// 使用 os.OpenFile、os.Create 或 os.Open 函数。
// 打开或创建文件
file, err := os.Create("output.txt")
if err != nil {
    fmt.Println("Error creating file:", err)
    return
}
defer file.Close()
```

#### 3.1 fmt.Fprint

​	将参数直接输出到文件。

```go
fmt.Fprint(file, "Hello", " ", "World!")
//file: Hello World!
```

#### 3.2 fmt.Fprintf

​	将参数按照格式化后，直接输出到文件。

```go
age := 26
name := "zhu"
fmt.Sprintf("Age:%d, Name:%s", age, name)
//file: Age:26, Name:zhu
```

#### 3.1 fmt.Fprintln

​	将参数之间加空格，末尾换行后，直接输出到文件。

```go
fmt.Fprint(file, "Hello", "World!")
//file: Hello World!
```

### 占位符

​	`	fmt` 包的格式化占位符种类繁多，以下是一些最常用的格式化占位符及其用法示例。

#### 1. 通用占位符

- `%v`：值的默认格式表示。
- `%+v`：结构体输出：key : value。
- `%#v`：Go 语法输出：struct { Name string }{Name:"Alice"}。
- `%T`：输出得是==值的类型==。
- `%%`：输出百分号（%）

#### 2. 类型占位符

- `%d`：十进制表示。
- `%f`：小数表示，如 `123.456`。

- `%s`：字符串表示法。

- `%t`：布尔值表示 `true` 或 `false`。