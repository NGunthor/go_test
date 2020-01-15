package reverse_llist

// ListNode provides interface for listNode struct
type ListNode interface {
	ReverseList() ListNode
	GetNext() ListNode
	SetNext(ListNode)
}

type listNode struct {
	val  int
	next ListNode
}

func (l *listNode) GetNext() ListNode {
	return l.next
}

func (l *listNode) SetNext(node ListNode) {
	l.next = node
}

func (l *listNode) ReverseList() ListNode {
	if l.GetNext() == nil {
		return l
	}
	cur := l.GetNext()
	var prev ListNode = l

	for cur != nil {
		next := cur.GetNext()
		cur.SetNext(prev)
		prev = cur
		cur = next
	}
	l.SetNext(nil)
	return prev
}

// NewListNode ...
func NewListNode(val int, next ListNode) ListNode {
	return &listNode{
		val:  val,
		next: next,
	}
}
