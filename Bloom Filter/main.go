package main

import (
	"fmt"
	// https://pkg.go.dev/github.com/bits-and-blooms/bloom/v3#section-readme
	"github.com/bits-and-blooms/bloom/v3"
)

func main() {
	// Initialize a Bloom Filter
	filter := bloom.New(20*1000*1000, 5) // Capacity: 20 million, 5 hash functions

	// Add a username to the Bloom Filter
	filter.AddString("john_doe")

	// Check if a username exists
	exists := filter.TestString("john_doe")
	fmt.Printf("Username 'john_doe' exists? %v\n", exists)

	// Check for a non-existent username
	exists = filter.TestString("jane_doe")
	fmt.Printf("Username 'jane_doe' exists? %v\n", exists)
}
