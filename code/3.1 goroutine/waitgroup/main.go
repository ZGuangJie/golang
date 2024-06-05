package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// 使用关键字 go 启动一个协程
	for i := 0; i < 5; i++ {
		// 进程启动时，增加一个
		wg.Add(1)
		go func() {
			// 进程结束时，减少一个
			defer wg.Done()
			fmt.Println("I'm gorotine...")
		}()
	}
	fmt.Println("I'm threading...")
	wg.Wait()
}
