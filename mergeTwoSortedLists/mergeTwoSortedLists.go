package mergeTwoSortedLists

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	output := make([]int, m + n)
	copy(output, nums1)
	output = output[:m]
	fmt.Println(output)
	for _, el := range nums2 {
		output = append(output, el)
	}
	sort.Ints(output)
	return output
}
