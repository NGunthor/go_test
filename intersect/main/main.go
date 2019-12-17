package main

import (
	"fmt"
	"github.com/NGunthor/go_test/intersect"
)

func main() {
	fmt.Println(intersect.Intersect([]int{4,9,5}, []int{9,4,9,8,4}))
	fmt.Println(intersect.Intersect([]int{1,2,2,1}, []int{2,2}))
}
