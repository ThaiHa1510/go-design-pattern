package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var n int32
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n *int32) {
			defer wg.Done()
			atomic.AddInt32(n, 1)
			atomic.
		}(&n)
	}
	wg.Wait()
	fmt.Println("n = ", n)
}
