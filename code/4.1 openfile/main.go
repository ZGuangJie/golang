package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 读取文件句柄
	file, err := os.ReadFile("data/test.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	// defer file.Close()
	// data, err := os.ReadFile(file)
	fmt.Printf("%v", string(file))
}
