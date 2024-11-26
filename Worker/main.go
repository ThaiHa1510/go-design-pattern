package main

import "fmt"

func main() {
	jobChan := make(chan int, 5)
	resultChan := make(chan int, 5)
	for i := 0; i < 3; i++ {
		go worker(i, jobChan, resultChan)
	}
	close(jobChan)
	for r := range resultChan {
		fmt.Println(r)
	}
}

// worker will range over the given channel and send each value to the result channel.
// Useful for running a goroutine that needs to process a stream of values from a channel.
func worker(i int, job <-chan int, result chan<- int) {
	for j := range job {
		result <- j
	}
}
