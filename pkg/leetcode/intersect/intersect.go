package intersect

// Numbers provides interface for numbers struct
type Numbers interface {
	Intersect() []int
}

type numbers struct {
	nums1 []int
	nums2 []int
}

// Intersect main algorithm function
func (n *numbers) Intersect() []int {
	m := map[int]int{}
	for _, el := range n.nums1 {
		m[el]++
	}
	out := make([]int, 0)
	for _, el := range n.nums2 {
		if m[el] > 0 {
			m[el]--
			out = append(out, el)
		}
	}
	return out
}

// NewNumbers ...
func NewNumbers(nums1, nums2 []int) Numbers {
	return &numbers{
		nums1: nums1,
		nums2: nums2,
	}
}
