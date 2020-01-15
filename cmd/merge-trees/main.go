package main

import (
	"fmt"
	merge_trees "github.com/NGunthor/go_test/pkg/leetcode/merge-trees"
)

func main() {
	tree1 := merge_trees.NewBinaryTree(1,2,3,4)
	tree2 := merge_trees.NewBinaryTree(1,2,3,4)
	result := merge_trees.NewTrees(tree1.GetHead(), tree2.GetHead()).MergeTrees()
	fmt.Println(result)
}
