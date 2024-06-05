package main

import (
	"fmt"
)

type Animals interface {
	say()
	eat()
}

type Dog struct {
}

func (dog Dog) say() {
	fmt.Println("wangwangwang")
}
func (dog Dog) eat() {
	fmt.Println("Only eat bone")
}
func (dog Dog) run() {
	fmt.Println("runrunrun")
}

type Cat struct {
}

func (cat *Cat) say() {
	fmt.Println("miaomiaomiao")
}
func (dog *Cat) eat() {
	fmt.Println("Only eat fish")
}

func main() {
	var X Animals
	// 定义一个Dog类，Dog类实现了Animals的全部方法，即实现了这个接口
	dog := Dog{}
	X = dog
	X.say()
	dog.run()
	// 类似的实现一个Cat类
	cat := Cat{}
	X = &cat
	X.say()

}
