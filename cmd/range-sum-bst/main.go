package main

import (
	"fmt"

	range_sum_bst "github.com/NGunthor/go_test/pkg/leetcode/range-sum-bst"
)

func main() {
	tree := range_sum_bst.NewBinaryTree(10, 5, 15, 3, 7, 18)
	fmt.Println(range_sum_bst.NewTree(tree.GetHead()).RangeSumBST(7, 15))
}
