package products

type product1 struct {
	name string
	age  int
}

// Use interacts with the product (implements User interface)
func (p *product1)Use() string {
	return p.name
}

// NewProduct1 ...
func NewProduct1(name string, age int) User {
	return &product1{
		name: name,
		age:  age,
	}
}
