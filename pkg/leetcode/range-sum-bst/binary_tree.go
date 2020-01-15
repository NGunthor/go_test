package range_sum_bst

// BinaryTree provides interface for binaryTree struct
type BinaryTree interface {
	GetHead() TreeNode
	SetHead(TreeNode)
}

type binaryTree struct {
	head TreeNode
}

func (b *binaryTree) GetHead() TreeNode {
	return b.head
}

func (b *binaryTree) SetHead(node TreeNode) {
	b.head = node
}

func (b *binaryTree) add(val int) {
	current := b.head
	for {
		if val > current.GetValue() {
			if current.GetRight() == nil {
				current.SetRight(NewTreeNode(val))
				break
			} else {
				current = current.GetRight()
			}
		} else if val < current.GetValue() {
			if current.GetLeft() == nil {
				current.SetLeft(NewTreeNode(val))
				break
			} else {
				current = current.GetLeft()
			}
		} else {
			continue
		}
	}

}

// NewBinaryTree creates binaryTree as BinaryTree interface
func NewBinaryTree(values ...int) BinaryTree {
	head := NewTreeNode(values[0])
	values = values[1:]
	tree := &binaryTree{head:head}
	for _, val := range values {
		tree.add(val)
	}
	return tree
}
