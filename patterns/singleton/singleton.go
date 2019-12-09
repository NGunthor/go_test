package singleton

import "sync"

type Singleton struct {
	Name string
	Age string
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func () {
		instance = &Singleton{}
	})
	return instance
}
