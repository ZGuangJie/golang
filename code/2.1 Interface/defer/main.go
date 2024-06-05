package main

import (
	"fmt"
)

type str struct {
	name string
}

func (s str) test() {
	fmt.Println(s.name, "Closed")
}

func main() {
	var s = [3]str{{"a"}, {"b"}, {"c"}}

	for _, i := range s {
		defer i.test()
	}
}
