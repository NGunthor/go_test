package factory_method

import (
	"fmt"

	products "github.com/NGunthor/go_test/pkg/patterns/factory-method/products"
)

type user interface {
	Use() string
}

// Creator provides interface for factory method
type Creator interface {
	CreateProduct(s string) user
}

type concreteCreator struct {
}

// CreateProduct creates new concrete product as factory method (implements Creator interface)
func (p * concreteCreator)CreateProduct(s string) user {
	var prod user

	switch s {
	case "1":
		prod = products.NewProduct1("qwerty", 10)
	case "2":
		prod = products.NewProduct2("asdfg", 15, true)
	case "3":
		prod = products.NewProduct3("zxcvb", 20, "qwertyasdfgzxcvb")
	default:
		{
			fmt.Println("Undefined type")
			return nil
		}
	}
	return prod
}

func NewCreator() Creator {
	return &concreteCreator{}
}