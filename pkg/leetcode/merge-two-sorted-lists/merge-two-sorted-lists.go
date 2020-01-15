package merge_two_sorted_lists

// Numbers provides interface for numbers struct
type Numbers interface {
	Merge() []int
}

type numbers struct {
	nums1 	[]int
	m		int
	nums2	[]int
	n		int
}

func (n *numbers)Merge() []int{
	a := make([]int, 0)
	n.nums1 = n.nums1[:n.m]
	n.nums2 = n.nums2[:n.n]
	for len(n.nums1) > 0 && len(n.nums2) > 0 {
		if n.nums1[0] > n.nums2[0] {
			a = append(a, n.nums2[0])
			n.nums2 = n.nums2[1:]
		} else {
			a = append(a, n.nums1[0])
			n.nums1 = n.nums1[1:]
		}

	}
	if len(n.nums1) == 0 {
		a = append(a, n.nums2...)
	} else {
		a = append(a, n.nums2...)
	}
	return a
}

// NewNumbers
func NewNumbers(nums1 []int, m int, nums2 []int, n int) Numbers {
	return &numbers{
		nums1: nums1,
		m:     m,
		nums2: nums2,
		n:     n,
	}
}
