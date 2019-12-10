package mergeTwoSortedLists

func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	a := make([]int, m)
	copy(a, nums1[:m])
	nums2 = nums2[:n]
	i := 0
	for len(a) > 0 && len(nums2) > 0 {
		if a[0] > nums2[0] {
			nums1[i] = nums2[0]
			nums2 = nums2[1:]
		} else {
			nums1[i] = a[0]
			a = a[1:]
		}
		i++
	}
	nums1 = nums1[:i]
	if len(a) == 0 {
		nums1 = append(nums1, nums2...)
	} else {
		nums1 = append(nums1, a...)
	}
	return nums1
}
