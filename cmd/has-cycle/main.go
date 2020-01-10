package main

import (
	"fmt"

	has_cycle "github.com/NGunthor/go_test/pkg/leetcode/has-cycle"
)

func main() {
	list := has_cycle.NewList(1,2,3,4,5)
	a := list.Get(4)
	b := list.Get(1)
	a.SetNext(b)
	fmt.Println(list.HasCycle())
}
