package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(runtime.NumCPU())
			wg.Done()
		}()
	}
	wg.Wait()
}


