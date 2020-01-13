package products

type product3 struct {
	name string
	age int
	disc string
}

// Use interacts with the product (implements User interface)
func (p *product3)Use() string {
	return p.name
}

// NewProduct3 ...
func NewProduct3(name string, age int, disc string) User {
	return &product3{
		name: name,
		age:  age,
		disc: disc,
	}
}