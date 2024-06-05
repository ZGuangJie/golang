package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.RWMutex

func read() {
	// 调用读锁，多个读之间不影响，写的时候才影响
	lock.RLock()
	fmt.Println("我在读呢")
	time.Sleep(time.Second)
	fmt.Println("读取完毕")
	lock.RUnlock()
}
func writer() {
	// 调用写锁，同时只有一个协程拥有写锁
	lock.Lock()
	fmt.Println("在写")
	time.Sleep(time.Second * 3)
	fmt.Println("写完了")
	lock.Unlock()
}
func main() {
	for i := 0; i < 5; i++ {
		go read()
	}
	go writer()
	time.Sleep(time.Second * 10)
}
