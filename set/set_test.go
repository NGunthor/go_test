package set

import (
	"reflect"
	"testing"
)

var addTest = []addTests{
	{
		[]int{1, 2, 2, 3, 4},
		[]int{1, 2, 3, 4},
	},
	{[]int{1, 3, 5, 6, 1, 2, 3, 100},
		[]int{1, 2, 3, 5, 6, 100},
	},
}

func TestSet_Add(t *testing.T) {
	for _, test := range addTest {
		s := NewSet()
		for _, element := range test.values {
			s.Add(element)
		}
		v := s.ToSlice()
		if !reflect.DeepEqual(v, test.result) {
			t.Error("For", test.values,
				"expected", test.result,
				"got", v)
		}
	}
}

var deleteTest = []deleteTests{
	{
		*NewSet(1, 2, 3, 4),
		[]int{1, 4, 6},
		[]int{2, 3},
	},
}

func TestSet_Delete(t *testing.T) {
	for _, test := range deleteTest {
		s := test.start
		for _, element := range test.values {
			s.Delete(element)
		}
		v := s.ToSlice()
		if !reflect.DeepEqual(v, test.result) {
			t.Error("For", test.values,
				"expected", test.result,
				"got", v)
		}
	}
}

var unionTest = []unionTests{
	{[]set{
		*NewSet(1, 2, 3, 4),
		*NewSet(2, 3, 4, 5),
		*NewSet(3, 4, 5, 6),
	},
		[]int{1, 2, 3, 4, 5, 6},
	},
}

func TestSet_Union(t *testing.T) {
	for _, test := range unionTest {
		var s set
		for _, element := range test.sets {
			s = s.Union(element)
		}
		v := s.ToSlice()
		if !reflect.DeepEqual(v, test.result) {
			t.Error("For", s,
				"expected", test.result,
				"got", v)
		}
	}
}

var intersectionTest = []intersectionTests{
	{*NewSet(1, 2, 3, 4),
		[]set{
			*NewSet(2, 3, 4, 5),
			*NewSet(3, 4, 5, 6),
		},
		[]int{3, 4},
	},
}

func TestSet_Intersection(t *testing.T) {
	for _, test := range intersectionTest {
		s := test.start
		for _, element := range test.sets {
			s = s.Intersection(element)
		}
		v := s.ToSlice()
		if !reflect.DeepEqual(v, test.result) {
			t.Error("For", s,
				"expected", test.result,
				"got", v)
		}
	}
}

var differenceTest = []differenceTests{
	{*NewSet(1, 2, 3, 4),
		*NewSet(3, 4),
		[]int{1, 2}},
}

func TestSet_Difference(t *testing.T) {
	for _, test := range differenceTest {
		a := test.a
		b := test.b
		v := a.Difference(b).ToSlice()
		if !reflect.DeepEqual(v, test.result) {
			t.Error("For", a, b,
				"expected", test.result,
				"got", v)
		}
	}
}

var subsetTest = []subsetTests{
	{*NewSet(1, 2, 3, 4, 5),
		*NewSet(2, 3),
		true},
}

func TestSet_IsSubset(t *testing.T) {
	for _, test := range subsetTest {
		v := test.set.IsSubset(test.subset)
		if !reflect.DeepEqual(v, test.result) {
			t.Error("For", test.set.ToSlice(), test.subset.ToSlice(),
				"expected", test.result,
				"got", v)
		}
	}
}
