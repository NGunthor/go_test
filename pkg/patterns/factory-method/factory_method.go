package factory_method

import (
	"fmt"
)

type Creator interface {
	CreateProduct(s string) Producter
}

type Producter interface {
	Use() string
}

func (p * ConcreteCreator)CreateProduct(s string) Producter {
	var prod Producter

	switch s {
	case "1":
		prod = &Product1{
			name: "asd",
			age:  10,
		}
	case "2":
		prod = &Product2{
			name:   "fdss",
			age:    34,
			broken: true,
		}
	case "3":
		prod = &Product3{
			name: "32423",
			age:  14,
			disc: "dafdfgad",
		}
	default:
		{
			fmt.Println("Undefined type")
			return nil
		}
	}
	return prod
}

type ConcreteCreator struct {

}

func NewCreator() Creator {
	return &ConcreteCreator{}
}

type Product1 struct {
	name string
	age int
}

func (p *Product1)Use() string {
	return p.name
}

type Product2 struct {
	name string
	age int
	broken bool
}

func (p *Product2)Use() string {
	return p.name
}

type Product3 struct {
	name string
	age int
	disc string
}

func (p *Product3)Use() string {
	return p.name
}