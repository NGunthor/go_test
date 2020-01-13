package adapter

// Driver provides driver's interface
type Driver interface {
	Travel(transport)
}

type driver struct {
}

// Travel performs traveling via Transport
func (d *driver) Travel(transport transport) {
	transport.Drive()
}

// NewDriver ...
func NewDriver() Driver {
	return &driver{}
}
