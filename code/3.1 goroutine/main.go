package main

import (
	"fmt"
)

func main() {
	// 使用关键字 go 启动一个协程
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("I'm gorotine...")
		}()
	}

	fmt.Println("I'm threading...")
}
