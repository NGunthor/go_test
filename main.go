package main

import (
	"fmt"
	"reflect"
	"tests/set"
)

func main() {
	s := set.NewSet()
	s.Add(5, 3, 7, 1, 8)
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(s.ToSlice())
}
