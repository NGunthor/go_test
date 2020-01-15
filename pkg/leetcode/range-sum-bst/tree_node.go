package range_sum_bst

// TreeNode provides interface for treeNode struct
type TreeNode interface {
	GetValue() int
	GetRight() TreeNode
	GetLeft() TreeNode
	SetRight(TreeNode)
	SetLeft(TreeNode)
}
type treeNode struct {
	val   int
	left  TreeNode
	right TreeNode
}

func (t *treeNode) GetValue() int {
	return t.val
}

func (t *treeNode) GetRight() TreeNode {
	return t.right
}

func (t *treeNode) GetLeft() TreeNode {
	return t.left
}

func (t *treeNode) SetLeft(treeNode TreeNode) {
	t.left = treeNode
}

func (t *treeNode) SetRight(treeNode TreeNode) {
	t.right = treeNode
}

// NewTreeNode ...
func NewTreeNode(val int) TreeNode {
	return &treeNode{
		val:   val,
		left:  nil,
		right: nil,
	}
}
