package testgo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("exceted:%#v, got: %#v", want, got)
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("枯藤老树昏鸦", "老")
	}
}

func ExampleSplit() {
	fmt.Println(Split("a:b:c", ":"))
	fmt.Println(Split("枯藤老树昏鸦", "老"))
	// Output:
	// [a b c]
	// [ 枯藤 树昏鸦]
}
