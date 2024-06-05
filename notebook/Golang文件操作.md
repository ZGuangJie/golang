### 1、Golang之文件

​		文件是保存数据的地方，是数据源的一种，比如经常使用的word、txt、excel、jpg...都是文件。文件最主要的作用就是保存数据，既可以是一张图片，也可使是视频、声音...

#### 1.1 文件的读取和关闭

​		使用os包下的Open()方法打开指定文件夹下的文件，打开后一般需要调用.Close()方法关闭文件。

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 返回的是指向文件的指针
	file, err := os.Open("data/test.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
    
	defer file.Close()
	fmt.Printf("%v", file)
}
/*
	使用os.Open()只是打开了文件，相当于箱子打开盖子，并没有拿出箱子里的东西，需要通过I/O流拿到里面的数据。
*/
```

### 2、I/O流

​		使用Open()打开的文件返回的File结构体并不能操作文件里的数据，只能看看文件名，文件类型一类的Info。通过io.ReadFile()读取的是文件中的数据，i/o流是程序和数据之间沟通的桥梁。

[![文件I/O流](https://cdn.jsdelivr.net/gh/ZGuangJie/GoPicture/golang/202406051221510.png)]()

##### 2.1 一次性读取所有数据

​		此种方法用于一次性读取文件里的所有数据，适合文件比较小的情况。

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 读取文件中的数据
	file, err := os.ReadFile("data/test.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	// defer file.Close()
	// data, err := os.ReadFile(file)
	fmt.Printf("%v", file)
}
/*
	使用io.ReadFile()打开文件不需要进行Open/Close操作，因为都封装在ReadFile内部了。
*/
```

##### 2.2 分批读取所有数据

​		bufio.NewReader()分批读取文件里的所有数据，默认每批次是4096byte。

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// 读取文件句柄
	file, err := os.Open("../data/test.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
        // 读取数据，使用换行符'\n'分割
		str, err := reader.ReadString('\n')
		fmt.Printf("%v", str)
		if err == io.EOF {
			break
		}
	}
}
/*
	实现分批次读取文件里的数据，然后使用换行符分割
	注意最后一行的末尾是io.EOF而不是\n
*/
```

