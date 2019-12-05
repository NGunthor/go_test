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
	{[]int{2, 7, 11, 15}, 9, []int{1, 1}},
	{[]int{4, 8, 0, 7, 2, 5}, 7, []int{2, 3}},
	{[]int{1, 4, 6, 7, 10, 47, 4, 6, 7, 10, 47, 4, 6, 7, 10, 47, 4, 6, 7, 10, 47, 4, 6, 7, 10, 47, 4, 6, 7, 10, 47, 4, 6, 7, 10, 47, 4, 6, 7, 10, 47, 4, 6, 7, 10, 47, 4, 6, 7, 10, 47, 4, 6, 7, 10, 47, 4, 6, 7, 10, 47, 4, 6, 7, 10, 47, 60}, 61, []int{0, 66}},
	{[]int{4, 8, 1, 4}, 8, []int{0, 3}},
}

func TestTwoSum(t *testing.T) {
	for _, test := range tests {
		v := TwoSum(test.values, test.target)
		if !reflect.DeepEqual(v, test.result) {
			t.Error(
				"For", test.values,
				"expected", test.result,
				"got", v)
		}

	}
}

func TestTwoSumFail(t *testing.T) {
	for _, test := range tests {
		v := TwoSum(test.values, test.target)
		if reflect.DeepEqual(v, test.result) {
			t.Log(
				"Success for", test.values,
				"expected", test.result,
				"got", v)
		} else {
			t.Failed()
			t.Log("Fail for", test.values,
				"expected", test.result,
				"got", v)
		}
	}
}