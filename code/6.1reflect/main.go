package main

import (
	"fmt"
	"reflect"
)

func reflectTest(v interface{}) {
	v_value := reflect.ValueOf(v)
	v_value.Elem().SetInt(2)

}
func main() {
	var i int = 3
	reflectTest(&i)
	fmt.Println(i)
}
