package main

import "fmt"

// Step 1 : Định nghĩa giao diện Prototype
type Prototype interface {
	Clone() Prototype
}

// Step2 : Triển khai các Concrete Protypes

type Circle struct {
	data string
}

func (c *Circle) Clone() Prototype {
	fmt.Println("Clone Circle")
	return &Circle{c.data}
}

func main() {
	// Tạo đối tượng Context
	circle := &Circle{"Circle"}

	// Sao chép đối tượng
	circle1 := circle.Clone().(*Circle)
	circle2 := circle.Clone().(*Circle)

	fmt.Println("Circle", circle.data)
	fmt.Println("Circle 1", circle1.data)
	fmt.Println("Circle 2", circle2.data)

	// Modifying data
	circle.data = "Modified"
	fmt.Println("Circle", circle.data)
	fmt.Println("Circle", circle1.data)
	fmt.Println("Circle", circle2.data)

}
