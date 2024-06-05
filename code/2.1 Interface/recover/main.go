package main

import (
	"errors"
	"fmt"
)

func main() {
	err := test()
	if err != nil {
		panic(err)
	}
	fmt.Println("我不一定能被执行...")
}

func test() error {

	num1 := 10
	num2 := 0
	if num2 == 0 {
		return errors.New("错了, 我不想执行了")
	}
	result := num1 / num2
	fmt.Println(result)
	return nil
}
