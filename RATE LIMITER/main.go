package main

import (
	"fmt"
	"time"
)

func main() {
	request := make(chan string, 5)
	for i := 0; i < 5; i++ {
		request <- fmt.Sprintf("Request %d", i)
	}
	close(request)
	limiter := time.Tick(2 * time.Second)
	for req := range request {
		<-limiter
		fmt.Println("Request", req, time.Now())
	}
}
