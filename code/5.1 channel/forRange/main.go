package main

import (
	"fmt"
	"sync"
	"time"
)

// 主线程等待协程执行完毕
var wg sync.WaitGroup

func writer(intChan chan int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		intChan <- i * 10
		fmt.Println("写入一个数据: ", i*10)
		time.Sleep(time.Second * 1)
	}
	// close(intChan)
}
func reader(intChan chan int) {
	defer wg.Done()
	for v := range intChan {
		fmt.Println("读了一个数据: ", v)
		time.Sleep(time.Second * 2)
	}
}
func main() {
	intChan := make(chan int, 10)
	wg.Add(2)

	go writer(intChan)
	go reader(intChan)

	wg.Wait()
}
