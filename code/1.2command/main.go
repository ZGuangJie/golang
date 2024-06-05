package command

import (
	"fmt"
	"os"
)

func Command() {
	var s string
	for index, value := range os.
		Args[0:] {
		fmt.Println(index, value)
	}
}
