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

func worker(i int, job <-chan int, result chan<- int) {
	for j := range job {
		result <- j
	}
}
