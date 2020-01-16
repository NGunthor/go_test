package product1

import (
	"fmt"
	"github.com/NGunthor/go_test/pkg/patterns/visitor/asdfasdf/product2"
)

type Product interface {
	Show()
}

type product struct {

}

func (p *product) Show() {
	fmt.Println(p)
	product2.NewProduct().Show()
}

func NewProduct() Product {
	return &product{}
}
