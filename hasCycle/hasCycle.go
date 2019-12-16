package hasCycle

//A struct that imagine a node of the List
type ListNode struct {
	Val  int
	Next *ListNode
}

//Method that runs the algorithm
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
