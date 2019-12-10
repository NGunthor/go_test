package factoryMethod

import (
	"reflect"
	"testing"
)

type Tests struct{
	value string
	result Producter
}

var tests = []Tests {
	{"1", NewCreator().CreateProduct("1")},
	{"2", NewCreator().CreateProduct("2")},
	{"3", NewCreator().CreateProduct("3")},
}

func TestNewCreator(t *testing.T) {
	for _, test := range tests {
		v := NewCreator().CreateProduct(test.value)
		if !reflect.DeepEqual(v, test.result) {
			t.Error("For", test.value,
				"expected", test.result,
				"got", v)
		}
	}
}
