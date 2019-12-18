package main

import "github.com/NGunthor/go_test/pkg/leetcode/merge-trees"

var t11 = merge_trees.TreeNode{
	Val:   1,
	Left:  nil,
	Right: nil,
}
var t12 = merge_trees.TreeNode{
	Val:   2,
	Left:  nil,
	Right: nil,
}
var t13 = merge_trees.TreeNode{
	Val:   3,
	Left:  nil,
	Right: nil,
}
var t14 = merge_trees.TreeNode{
	Val:   5,
	Left:  nil,
	Right: nil,
}

var t21 = merge_trees.TreeNode{
	Val:   1,
	Left:  nil,
	Right: nil,
}
var t22 = merge_trees.TreeNode{
	Val:   2,
	Left:  nil,
	Right: nil,
}
var t23 = merge_trees.TreeNode{
	Val:   3,
	Left:  nil,
	Right: nil,
}
var t24 = merge_trees.TreeNode{
	Val:   4,
	Left:  nil,
	Right: nil,
}
var t25 = merge_trees.TreeNode{
	Val:   7,
	Left:  nil,
	Right: nil,
}

func main() {
	t11.Left = &t12
	t12.Left = &t13

	t21.Right = &t22
	t22.Right = &t23

	merge_trees.MergeTrees(&t11, &t21)
}
