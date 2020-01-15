package main

import (
	"fmt"
	two_sums "github.com/NGunthor/go_test/pkg/leetcode/two-sums"
)

func main() {
	result := two_sums.NewNumbers([]int{1,2,3,4,5,6}, 10)
	fmt.Println(result.TwoSum())
}
