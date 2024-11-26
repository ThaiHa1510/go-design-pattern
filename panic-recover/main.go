package main

import (
	"fmt"
)

func mayPanic() {
	// This function deliberately causes a panic
	panic("a problem occurred")
}

func main() {
	// Defer a function to recover from a panic
	defer func() {
		if r := recover(); r != nil {
			// If there is a panic, this block will execute
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// This function is expected to panic
	mayPanic()

	mayPanic()
	// This line will only be reached if the panic was recovered
	fmt.Println("The main function continues after recovery.")
}
