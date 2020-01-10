package main

import w "github.com/NGunthor/go_test/pkg/leetcode/worker-pool"

func main() {
	f := w.NewForeman(40, 5)
	f.Start()
}
