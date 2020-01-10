package has_cycle

// ListNode provides interface for ListNode struct
type ListNode interface {
	SetNext(*listNode)
}

type listNode struct {
	val  int
	next *listNode
}

// SetNext sets $node to current listNode
func (l *listNode) SetNext(node *listNode) {
	l.next = node
}

// NewListNode ...
func NewListNode(val int) *listNode {
	return &listNode{
		val:  val,
		next: nil,
	}
}
