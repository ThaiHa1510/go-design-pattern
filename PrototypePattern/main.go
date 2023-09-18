package main

import "fmt"

// Step 1: Định nghĩa giao diện Prototype
type Prototype interface {
	Clone() Prototype
}

// Step 2: Triển khai các Concrete Prototypes
type ConcretePrototype struct {
	data string
}

func (c *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{data: c.data}
}

// Step 3: Sử dụng Prototype
func main() {
	original := &ConcretePrototype{data: "Original data"}
	copy1 := original.Clone().(*ConcretePrototype)
	copy2 := original.Clone().(*ConcretePrototype)

	fmt.Println("Original data:", original.data)
	fmt.Println("Copy 1 data:", copy1.data)
	fmt.Println("Copy 2 data:", copy2.data)

	// Thay đổi dữ liệu của copy1
	copy1.data = "Modified data"

	fmt.Println("Original data:", original.data)
	fmt.Println("Copy 1 data:", copy1.data)
	fmt.Println("Copy 2 data:", copy2.data)
}
