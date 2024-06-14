package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 10)

	for i := 1; i <= 10; i++ {
		intChan <- i * 10
		fmt.Println("写入一个数据: ", i*10)
		time.Sleep(time.Second * 1)
	}
	select {}
}
