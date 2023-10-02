package main

import (
	"fmt"
	"sync"
)

func GetTotal(waitGroup *sync.WaitGroup, data []interface{}) {
	defer waitGroup.Done()
	data = append(data, "hello")
	fmt.Println("GetTotal")
}
func main() {
	var wg sync.WaitGroup
	data := []interface{}{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go GetTotal(&wg, data)
	}
	wg.Wait()
}
