package adapter

import "github.com/NGunthor/go_test/pkg/patterns/adapter/transports"

type animalToTransportAdapter struct {
	animal transports.Animal
}

// Drive performs drive work for Transport
func (h *animalToTransportAdapter) Drive() {
	h.animal.Move()
}

// NewAnimalToTransportAdapter ...
func NewAnimalToTransportAdapter(animal transports.Animal) *animalToTransportAdapter {
	return &animalToTransportAdapter{animal: animal}
}
