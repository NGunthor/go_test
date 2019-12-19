package adapter

import "github.com/NGunthor/go_test/pkg/patterns/adapter/transports"

type driver struct {
}

func (d *driver) Travel(transport transports.Transport) {
	transport.Drive()
}

func NewDriver() *driver {
	return &driver{}
}
