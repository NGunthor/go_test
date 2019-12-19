package adapter

import "github.com/NGunthor/go_test/pkg/patterns/adapter/transports"

type driver struct {
}

// Travel performs traveling via Transport
func (d *driver) Travel(transport transports.Transport) {
	transport.Drive()
}

// NewDriver ...
func NewDriver() *driver {
	return &driver{}
}
