package main

import (
	"fmt"

	"github.com/NGunthor/go_test/pkg/leetcode/intersect"
)

func main() {
	numbers := intersect.NewNumbers([]int{1,2,3,4}, []int{2,3})
	fmt.Println(numbers.Intersect())
}
