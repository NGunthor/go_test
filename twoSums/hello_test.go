package main

import (
	"reflect"
	"testing"
)

type testStruct struct{
	values []int
	target int
	result []int
}

var tests = []testStruct {
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{4, 8, 0, 7, 2, 5}, 7, []int{2, 3}},
	{[]int{1, 4, 6, 7, 10, 47, 4, 6, 7, 10, 47, 4, 6, 7, 10, 47, 4, 6, 7, 10, 47, 4, 6, 7, 10, 47, 4, 6, 7, 10, 47, 4, 6, 7, 10, 47, 4, 6, 7, 10, 47, 4, 6, 7, 10, 47, 4, 6, 7, 10, 47, 4, 6, 7, 10, 47, 4, 6, 7, 10, 47, 4, 6, 7, 10, 47, 60}, 61, []int{0, 66}},
	{[]int{4, 8, 1, 4}, 8, []int{0, 3}},
}

func TestTwoSum(t *testing.T) {
	for _, el := range tests {
		v := TwoSum(el.values, el.target)
		if !reflect.DeepEqual(v, el.result) {
			t.Error(
				"For", el.values,
				"expected", el.result,
				"got", v)
		}
	}
}
