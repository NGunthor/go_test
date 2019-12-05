package main

import "testing"

type testStruct struct {
	values [2]string
	result bool
}

var failTests = []testStruct {
	{[2]string{"aba", "bab"}, false},
	{[2]string{"foo", "bar"}, true},
}

var successTests = []testStruct {
	{[2]string{"ab", "aa"}, false},
	{[2]string{"aab", "aba"}, false},
	{[2]string{"egg", "add"}, true},
	{[2]string{"paper", "title"}, true},
}

func TestIsIsomorphicSuccess(t *testing.T) {
	for _, test := range successTests {
		v := isIsomorphic(test.values[0], test.values[1])
		if v != test.result {
			t.Error("For", test.values,
				"expected", test.result,
				"got", v)
		}
	}
}

func TestIsIsomorphicFail(t *testing.T) {
	for _, test := range failTests {
		v := isIsomorphic(test.values[0], test.values[1])
		if v == test.result {
			t.Fail()
			t.Log("For", test.values,
				"expected", test.result,
				"got", v)
		}
	}
}
