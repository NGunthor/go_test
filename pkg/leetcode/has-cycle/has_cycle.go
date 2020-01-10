package has_cycle

// List provides interface for list struct
type List interface {
	HasCycle() bool
	Get(n int) *listNode
}

type list struct {
	head *listNode
	size int
}

//Method that runs the algorithm
func (l *list)HasCycle() bool {
	m := map[*listNode]bool{}

	for l.head != nil {

		if m[l.head] == false {
			m[l.head] = true
			l.head = l.head.next
			continue
		}
		return true
	}
	return false
}

// Get gets listNode with the same value
func (l *list)Get(n int) *listNode {
	cur := l.head
	for cur != nil {
		if cur.val == n {
			return cur
		}
		cur = cur.next
	}
	return nil
}

// NewList ...
func NewList(values ...int) List {
	var head *listNode
	var prev *listNode
	size := 0
	for i, v := range values {
		if i == 0 {
			head = NewListNode(v)
			prev = head
			size++
			continue
		}
		prev.next = NewListNode(v)
		prev = prev.next
		size++
	}
	return &list{
		head: head,
		size: size,
	}
}
