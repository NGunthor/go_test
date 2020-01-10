package mac_factory

import "fmt"

type Checkbox interface {
	Paint()
}

type macCheckbox struct {
}

func (m *macCheckbox) Paint() {
	fmt.Println("Mac Checkbox painted")
}
