package products

type product2 struct {
	name string
	age int
	broken bool
}

// Use interacts with the product (implements User interface2)
func (p *product2)Use() string {
	return p.name
}

// NewProduct2 ...
func NewProduct2(name string, age int, broken bool) User {
	return &product2{
		name:   name,
		age:    age,
		broken: broken,
	}
}
