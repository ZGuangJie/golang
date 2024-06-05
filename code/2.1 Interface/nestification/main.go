package main

import "fmt"

type Animals interface {
	eat()
	say()
}

type Communication struct {
}

func (say Communication) say() {
	fmt.Println("Wangwangwang")
}

type Cat struct {
	Communication
}

// Cat实际上只实现了Animals接口的一个函数，
// 另外一个是从Communication继承过来的
func (cat Cat) eat() {
	fmt.Println("Only eat fish")
}

func main() {
	// var f float32
	var X Animals = Cat{}
	X.say()
}
