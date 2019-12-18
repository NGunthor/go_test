package singleton

import (
	"strconv"
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	a := GetInstance()
	b := GetInstance()
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(a Singleton, b Singleton, i int) {
			defer wg.Done()
			a.Name = strconv.Itoa(i)
			i++
			b.Name = strconv.Itoa(i)
			b.Age = strconv.Itoa(i)
			if a.Name != b.Name && b.Age == a.Age {
				t.Error("false")
			}
		}(*a, *b, i)
	}
	wg.Wait()
}
