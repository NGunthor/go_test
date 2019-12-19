package add_two_numbers

// ListNode represents a node of List
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	rem := 0
	head := &ListNode{}
	cur := head
	var last *ListNode
	for l1 != nil || l2 != nil {
		v1 := 0
		v2 := 0
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}

		result := v1 + v2 + rem
		rem = result / 10

		cur.Val = result % 10
		cur.Next = &ListNode{}
		last = cur
		cur = cur.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if rem == 0 {
		last.Next = nil
	} else {
		last.Next.Val = rem
	}
	return head
}

// MakeList makes list
func MakeList(ints ...int) *ListNode {
	head := &ListNode{}
	cur := head
	var last *ListNode
	for _, num := range ints {
		cur.Val = num
		cur.Next = &ListNode{}
		last = cur
		cur = cur.Next
	}
	last.Next = nil
	return head
}
