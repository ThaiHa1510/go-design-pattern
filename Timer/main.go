package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	tick := time.Tick(1 * time.Second)
	for {
		select {
		case <-done:
			fmt.Println("done")
			return
		case t := <-tick:
			fmt.Println("Tick at ,", t)
		}
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Run after")
	done <- true
}
