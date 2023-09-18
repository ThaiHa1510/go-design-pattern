package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Get Time excuted syncMap
	startTime := time.Now()
	syncMap()
	elapsed := time.Since(startTime)
	// print elapsed time
	fmt.Printf("Elapsed time: %s", elapsed)

}
func syncMap() {
	var wg sync.WaitGroup
	m := sync.Map{}

	// Sử dụng Goroutine để ghi dữ liệu vào sync.Map
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n)
			value := fmt.Sprintf("value%d", n)
			m.Store(key, value)
		}(i)
	}
	time.Sleep(5 * time.Second)
	// Sử dụng Goroutine để đọc dữ liệu từ sync.Map
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n)
			result, ok := m.Load(key)
			if ok {
				fmt.Printf("Value of %s: %s\n", key, result)
			} else {
				fmt.Printf("Key %s not found\n", key)
			}
		}(i)
	}

	wg.Wait()
}
