package main

import (
	"fmt"
)

func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producing %d\n", i)
		ch <- i
		// time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for val := range ch {
		fmt.Printf("Consuming %d\n", val)
		// time.Sleep(time.Millisecond * 800)
	}
}

func main() {
	ch := make(chan int, 2) // Buffered channel with capacity 2

	go producer(ch)
	consumer(ch)

	fmt.Println("Done")
}
