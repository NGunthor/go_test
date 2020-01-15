package merge_trees

// BinaryTreeNode provides interface for binaryTree struct
type BinaryTreeNode interface {
	SetValue(int)
	SetLeft(BinaryTreeNode)
	SetRight(BinaryTreeNode)
	GetValue() int
	GetRight() BinaryTreeNode
	GetLeft() BinaryTreeNode
}

type binaryTreeNode struct {
	val   int
	left  BinaryTreeNode
	right BinaryTreeNode
}

func (b *binaryTreeNode) SetLeft(node BinaryTreeNode) {
	b.left = node
}

func (b *binaryTreeNode) SetRight(node BinaryTreeNode) {
	b.right = node
}

func (b *binaryTreeNode) SetValue(val int) {
	b.val = val
}

func (b *binaryTreeNode) GetValue() int {
	return b.val
}

func (b *binaryTreeNode) GetRight() BinaryTreeNode {
	return b.right
}

func (b *binaryTreeNode) GetLeft() BinaryTreeNode {
	return b.left
}

// NewBinaryTreeNode ...
func NewBinaryTreeNode(val int) BinaryTreeNode {
	return &binaryTreeNode{
		val:   val,
		left:  nil,
		right: nil,
	}
}
