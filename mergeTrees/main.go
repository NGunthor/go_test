package mergeTrees

var t11 = TreeNode{
	Val:   1,
	Left:  nil,
	Right: nil,
}
var t12 = TreeNode{
	Val:   2,
	Left:  nil,
	Right: nil,
}
var t13 = TreeNode{
	Val:   3,
	Left:  nil,
	Right: nil,
}
var t14 = TreeNode{
	Val:   5,
	Left:  nil,
	Right: nil,
}

var t21 = TreeNode{
	Val:   1,
	Left:  nil,
	Right: nil,
}
var t22 = TreeNode{
	Val:   2,
	Left:  nil,
	Right: nil,
}
var t23 = TreeNode{
	Val:   3,
	Left:  nil,
	Right: nil,
}
var t24 = TreeNode{
	Val:   4,
	Left:  nil,
	Right: nil,
}
var t25 = TreeNode{
	Val:   7,
	Left:  nil,
	Right: nil,
}

func main() {
	t11.Left = &t12
	t12.Left = &t13

	t21.Right = &t22
	t22.Right = &t23

	MergeTrees(&t11, &t21)
}
