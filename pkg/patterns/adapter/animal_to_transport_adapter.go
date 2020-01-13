package adapter

type transport interface {
	Drive()
}

type animal interface {
	Move()
}

type animalToTransportAdapter struct {
	animal animal
}

// Drive performs drive work for Transport
func (h *animalToTransportAdapter) Drive() {
	h.animal.Move()
}

// NewAnimalToTransportAdapter ...
func NewAnimalToTransportAdapter(animal animal) transport {
	return &animalToTransportAdapter{animal: animal}
}
