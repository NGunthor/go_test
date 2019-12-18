package merge_trees

//A struct that imagine a node of the Tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	outT := TreeNode{}

	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	outT.Val = t1.Val + t2.Val
	outT.Left = MergeTrees(t1.Left, t2.Left)
	outT.Right = MergeTrees(t1.Right, t2.Right)
	return &outT
}
