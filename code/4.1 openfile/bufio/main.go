package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// 读取文件句柄
	file, err := os.Open("../data/test.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		fmt.Printf("%v,err:%v\n", str, err)
		if err == io.EOF {
			break
		}
	}
}
