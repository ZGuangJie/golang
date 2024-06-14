package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 10)

	go writer(intChan)

	for v := range intChan {
		fmt.Printf("输出: %v\n", v)
		time.Sleep(time.Second)
	}

}
func writer(intChan chan int) {
	for {
		select {
		case intChan <- 1:
			fmt.Println("Write 1...")
		default:
			fmt.Println("Channel full")
		}
		time.Sleep(time.Millisecond * 500)
	}
}
