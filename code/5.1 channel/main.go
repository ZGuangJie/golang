package main

import "fmt"

func main() {
	intChan := make(chan int, 2)
	intChan <- 10
	fmt.Printf("%#v", <-intChan)
	close(intChan)
}
