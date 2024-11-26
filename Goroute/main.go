package main

import (
	"context"
	"fmt"
	"time"
)

// func worker(id int, jobs <-chan int, results chan<- int) {
// 	for j := range jobs {
// 		fmt.Printf("Worker %d started job %d\n", id, j)
// 		time.Sleep(time.Second) // Simulate work
// 		fmt.Printf("Worker %d finished job %d\n", id, j)
// 		results <- j * 2
// 	}
// }

func main() {
	done := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	// jobs := make(chan int, 5)
	results := make(chan int, 5)
	go producer(results, cancel)
	go consumer(results, done, ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context done")
			return
		case <-done:
			fmt.Println("singal done chanel")
			return
		default:
			fmt.Println("default")
			time.Sleep(time.Second)
		}
	}
}

func consumer(ch chan int, done chan<- bool, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Receive done\n")
			return
		case <-ch:
			fmt.Printf("Consuming %d\n", <-ch)
			time.Sleep(time.Millisecond * 800)
		}
	}
	done <- true
}

func producer(ch chan int, cancel context.CancelFunc) {
	for i := 1; i <= 15; i++ {
		fmt.Printf("Producing %d\n", i)
		ch <- i
		time.Sleep(time.Millisecond * 500)
		if i == 10 {
			fmt.Printf("Emit done %d\n", i)
			cancel()
		}
	}
	close(ch)
}
