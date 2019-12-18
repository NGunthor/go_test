package intersect

// Intersect main algorithm function
func Intersect(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	for _, el := range nums1 {
		m[el]++
	}
	out := make([]int, 0)
	for _, el := range nums2 {
		if m[el] > 0 {
			m[el]--
			out = append(out, el)
		}
	}
	return out
}
