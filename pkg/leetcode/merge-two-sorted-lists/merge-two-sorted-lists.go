package merge_two_sorted_lists

func Merge(nums1 []int, m int, nums2 []int, n int) []int{
	a := make([]int, 0)
	nums1 = nums1[:m]
	nums2 = nums2[:n]
	for len(nums1) > 0 && len(nums2) > 0 {
		if nums1[0] > nums2[0] {
			a = append(a, nums2[0])
			nums2 = nums2[1:]
		} else {
			a = append(a, nums1[0])
			nums1 = nums1[1:]
		}

	}
	if len(nums1) == 0 {
		a = append(a, nums2...)
	} else {
		a = append(a, nums2...)
	}
	return a
}
