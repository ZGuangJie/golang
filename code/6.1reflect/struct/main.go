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
