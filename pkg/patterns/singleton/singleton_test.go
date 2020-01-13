package singleton

import (
	"strconv"
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	a := NewSingleton("qwert", 10)
	b := NewSingleton("qwer", 23)
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(a Singleton, b Singleton, i int) {
			defer wg.Done()
			a.SetName(strconv.Itoa(i))
			i++
			b.SetName(strconv.Itoa(i))
			b.SetAge(i)
			if a.GetName() != b.GetName() && b.GetAge() == a.GetAge() {
				t.Error("false")
			}
		}(a, b, i)
	}
	wg.Wait()
}
