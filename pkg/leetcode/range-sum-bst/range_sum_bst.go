package range_sum_bst

type Tree interface {
	RangeSumBST(int, int) int
}

type tree struct {
	head TreeNode
}

func (t *tree) RangeSumBST(L int, R int) int {
	if t.head == nil {
		return 0
	}

	var result int
	if t.head.GetValue() < L {
		result += NewTree(t.head.GetRight()).RangeSumBST(L, R)
	}
	if t.head.GetValue() > R {
		result += NewTree(t.head.GetLeft()).RangeSumBST(L, R)
	}
	if t.head.GetValue() >= L && t.head.GetValue() <= R {
		result += t.head.GetValue()
		if t.head.GetValue() != R {
			result += NewTree(t.head.GetRight()).RangeSumBST(L, R)
		}
		if t.head.GetValue() != L {
			result += NewTree(t.head.GetLeft()).RangeSumBST(L, R)
		}
	}
	return result
}

func NewTree(node TreeNode) Tree {
	return &tree{head:node}
}
