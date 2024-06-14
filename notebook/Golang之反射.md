### Golang之反射

​		在编写代码时，可能需要一个函数能够处理一类并不满足普通公共接口的类型的值，也可能是因为他们并没有确定的表示方式，或者在设计函数时，某些类型还不存在。[`reflect`](https://golang.org/pkg/reflect/) 实现了运行时的反射能力，能够让程序操作不同类型的对象[1](https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-reflect/#fn:1)。

​		**反射**是指在程序**`运行期`**对程序本身进行**访问和修改**的能力.

​		通过**空接口**可以被任意的类型的变量赋值的特性，设计一种操纵任意类型的方式，叫做`反射`。

#### 1、反射可以用来做什么？

1. 反射可以在**运行时**动态获取变量的各种信息，比如变量的类型，类别等信息
2. 如果是结构体变量，还可以获取到结构体本身的信息（包括结构体的字段、方法）
3. **通过反射，可以修改变量的值，可以调用关联的方法**。
4. 反射在reflect包里



#### 2、两个重要的函数

​	[`reflect`](https://golang.org/pkg/reflect/) 实现了运行时的反射能力，能够让程序操作不同类型的对象[1](https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-reflect/#fn:1)。反射包中有两对非常重要的函数和类型。

##### [`reflect.TypeOf`](https://draveness.me/golang/tree/reflect.TypeOf) and [`reflect.ValueOf`](https://draveness.me/golang/tree/reflect.ValueOf)，reflect.TypeOf 显示的是传入变量的类型，返回的是reflect.Type类型，要区别于普通的内置类型。

```go
package main

import (
	"fmt"
	"reflect"
)

func reflectTest(v interface{}) {
	v_type := reflect.TypeOf(v)
	v_value := reflect.ValueOf(v)
	// 这个类型是 reflect.Type
	fmt.Println("类型：", v_type)
	// 这个类型是 reflect.Value
	fmt.Println("值：", v_value)
}
func main() {
	var i int = 1
	reflectTest(i)
}
/*
	这样就可以通过reflectTest函数操作不累类型的值
*/
```

#### 3、使用反射操作内置的基本类型

##### 3.1 通过 断言 操作基本数据类型

​		reflect.Value可通过内置方法.Interface()，转换到接口，然后再通过类型断言，转换至原来类型的变量。

```go
package main

import (
	"fmt"
	"reflect"
)
func reflectTest(v interface{}) {
	v_value := reflect.ValueOf(v)
	v = v_value.Interface()
    // 转换到原本数据类型后正常使用
	if value, ok := v.(int); ok {
		value += 2
		fmt.Printf("将值增加到: %d", value)
	}
}
func main() {
	var i int = 3
	reflectTest(i)
}
/*
	
*/
```

##### 3.2 通过reflect内置的方法操作变量

​		通过内置方法转化为对应的类型。

```go
package main

import (
	"fmt"
	"reflect"
)

func reflectTest(v interface{}) {
	v_value := reflect.ValueOf(v)
	m := v_value.Int() + 2
	fmt.Println(m)
}
func main() {
	var i int = 3
	reflectTest(i)
}
```

​		通过指针修改传入变量的值。

```go
package main

import (
	"fmt"
	"reflect"
)

func reflectTest(v interface{}) {
	v_value := reflect.ValueOf(v)
	v_value.Elem().SetInt(2)

}
func main() {
	var i int = 3
	reflectTest(&i)
	fmt.Println(i)
}
```



#### 4、使用反射操作 结构体 对应的属性和方法

![结构体变量操作流程图](https://cdn.jsdelivr.net/gh/ZGuangJie/GoPicture/golang/202406081700974.png)

​		使用反射操作 结构体变量 也可以像操作基本数据类型一样，通过断言来将其转化为 基本数据类型，这里不在赘述。下面主要说的是通过内置的一些函数，操作结构体的**属性**和**绑定的方法**。（常用）

##### 4.1、操作属性

​		通过.NumFiled() 获取能操作的字段数，然后给 Filed() 方法传入索引，以访问参数。

```go
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func (s *Student) GetName() {
	fmt.Print(s.Name)
}

func (s *Student) AddAge(age int) bool {
	var flag bool
	defer func() {
		err := recover()
		fmt.Println("Fail add age:", err)
		flag = false
	}()
	s.Age += age
	flag = true
	return flag
}
func reflectTest(v interface{}) {
	v_value := reflect.ValueOf(v)
	fmt.Printf("%#v", v_value)
	// 使用NumField获取能够访问的字段个数
	num := v_value.NumField()
	// 使用Field输出对应字段的值
	for i := 0; i < num; i++ {
		fmt.Println(v_value.Field(i))
	}
}
func main() {
	var stu = Student{
		Name: "zhu",
		Age:  26,
	}
	reflectTest(stu)
}
```



##### 4.1、操作方法

​		通过.NumMethod() 方法获取能操作的方法个数（只能获得大写字符开头的方法），然后给 Method()  方法传入索引，最后通过Call(nil)方法调用。

​		`注意:`

1. 结构体绑定的方法中，是以ASCII码的大小排序的。
2. 使用Call()，传递参数要传递reflect.Value类型的参数。

```go
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func (s Student) GetName() {
	fmt.Println(s.Name)
	fmt.Println(s.Age)
}

func (s *Student) AddAge(age int) bool {
	var flag bool
	defer func() {
		err := recover()
		fmt.Println("Fail add age:", err)
		flag = false
	}()
	s.Age += age
	flag = true
	return flag
}

func reflectTest(v interface{}) {
	v_value := reflect.ValueOf(v)
	// fmt.Printf("%#v", v_value)
	// 使用NumMethod获取能够访问的方法的个数
	fmt.Println(v_value.NumMethod())
	// 使用Method输出对应字段的值
	fmt.Println(v_value.Method(0).Call(nil))

}
func main() {
	var stu = Student{
		Name: "zhu",
		Age:  26,
	}
	reflectTest(stu)
}
```

