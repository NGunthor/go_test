package main

import (
	"fmt"

	intersect2 "github.com/NGunthor/go_test/pkg/leetcode/intersect"
)

func main() {
	fmt.Println(intersect2.Intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	fmt.Println(intersect2.Intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}
