package merge_trees

// BinaryTree provides interface for binaryTree struct
type BinaryTree interface {
	GetHead() BinaryTreeNode
	SetHead(BinaryTreeNode)
}

type binaryTree struct {
	head BinaryTreeNode
}

func (b *binaryTree) GetHead() BinaryTreeNode {
	return b.head
}

func (b *binaryTree) SetHead(node BinaryTreeNode) {
	b.head = node
}

func (b *binaryTree) add(val int) {
	current := b.head
	for {
		if val > current.GetValue() {
			if current.GetRight() == nil {
				current.SetRight(NewBinaryTreeNode(val))
				break
			} else {
				current = current.GetRight()
			}
		} else if val < current.GetValue() {
			if current.GetLeft() == nil {
				current.SetLeft(NewBinaryTreeNode(val))
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
	head := NewBinaryTreeNode(values[0])
	values = values[1:]
	tree := &binaryTree{head:head}
	for _, val := range values {
		tree.add(val)
	}
	return tree
}
