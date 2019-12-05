package main

import "testing"

type testStruct struct {
	values [2]string
	result bool
}

var tests = []testStruct {
	{[2]string{"ab", "aa"}, false},
	{[2]string{"aab", "aba"}, false},
	{[2]string{"aba", "bab"}, false},
	{[2]string{"egg", "add"}, true},
	{[2]string{"foo", "bar"}, true},
	{[2]string{"paper", "title"}, true},
}

func TestIsIsomorphic(t *testing.T) {
	for _, test := range tests {
		v := isIsomorphic(test.values[0], test.values[1])
		if v == test.result {
			t.Log("Success for", test.values,
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
