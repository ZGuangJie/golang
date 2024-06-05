package main

import "fmt"

func add(a int) func() {
	var temp int
	return func() {
		temp += a
		fmt.Println(temp)
	}

}

func main() {

	c := add(1)
	c() // 第一次调用输出1
	c() // 第二次调用输出2
}
