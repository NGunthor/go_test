package hasCycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	m := map[*ListNode]bool{}

	for head != nil{

		if m[head] == false {
			m[head] = true
			head = head.Next
			continue
		}
		return true
	}
	return false
}
