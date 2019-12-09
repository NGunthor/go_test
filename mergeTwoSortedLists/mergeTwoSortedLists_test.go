package mergeTwoSortedLists

import (
	"reflect"
	"testing"
)

type Tests struct {
	nums1 []int
	m int
	nums2 []int
	n int
	result []int
}

var tests = []Tests{
	{
		nums1:  []int{1, 3,3,4,5,8, 0, 0, 0, 0},
		m:      6,
		nums2:  []int{4,5,6,10},
		n:      4,
		result: []int{1, 3, 3, 4, 4, 5, 5, 6, 8, 10},
	},
}

func TestMerge(t *testing.T) {
	for _, test := range tests {
		v := Merge(test.nums1, test.m, test.nums2, test.n)
		if !reflect.DeepEqual(v, test.result) {
			t.Error("For", test.nums1, test.nums2,
				"expected", test.result,
				"got", v)
		}
	}
}
