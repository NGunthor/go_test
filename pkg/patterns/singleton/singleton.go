package singleton

import "sync"

var (
	instance Singleton
	once sync.Once
)

// Singleton provides singleton's interface
type Singleton interface {
	DoWork()
	SetName(name string)
	SetAge(age int)
	GetName() string
	GetAge() int
}

type singleton struct {
	name string
	age int
}

// DoWork does a work via singleton (implements Singleton interface)
func (s *singleton) DoWork() {
	// ...
}

// SetName ...
func (s *singleton) SetName(name string) {
	s.name = name
}

// SetAge ...
func (s *singleton) SetAge(age int) {
	s.age = age
}

// GetName ...
func (s *singleton) GetName() string {
	return s.name
}

// GetAge ...
func (s *singleton) GetAge() int {
	return s.age
}

// NewSingleton ...
func NewSingleton(name string, age int) Singleton {
	once.Do(func () {
		instance = &singleton{
			name: name,
			age:  age,
		}
	})
	return instance
}
