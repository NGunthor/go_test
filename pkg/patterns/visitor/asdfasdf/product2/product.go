package product2

import (
	"fmt"
	"github.com/NGunthor/go_test/pkg/patterns/visitor/asdfasdf/product1"
)

type Product interface{
	Show()
}

type product struct {

}

func (p *product) Show() {
	fmt.Println(p)
	product1.NewProduct().Show()
}

func NewProduct() Product {
	return &product{}
}
