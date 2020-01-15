package merge_trees

// Trees provides interface for trees struct
type Trees interface {
	MergeTrees() BinaryTreeNode
}

type trees struct {
	t1 BinaryTreeNode
	t2 BinaryTreeNode
}

func (t *trees) MergeTrees() BinaryTreeNode {
	if t.t1 == nil {
		return t.t2
	}
	if t.t2 == nil {
		return t.t1
	}
	outT := NewBinaryTreeNode(t.t1.GetValue() + t.t2.GetValue())
	outT.SetLeft(NewTrees(t.t1.GetLeft(), t.t2.GetLeft()).MergeTrees())
	outT.SetRight(NewTrees(t.t1.GetRight(), t.t2.GetRight()).MergeTrees())
	return outT
}

// NewTrees ...
func NewTrees(t1, t2 BinaryTreeNode) Trees {
	return &trees{
		t1: t1,
		t2: t2,
	}
}
