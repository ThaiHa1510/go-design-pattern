package main

import "fmt"

func main() {
	// init array string
	arrStr := []string{"This", "is", "a", "test", "multi", "line"}

	fmt.Println("Original Array: ", arrStr)
	slice := arrStr[1:3]
	fmt.Println("Slice: ", slice)

	// print lengh slice
	fmt.Println("Length Slice: ", len(slice))
	// print capacity slice
	fmt.Println("Capacity Slice: ", cap(slice))
	// In cap array
	fmt.Println("Capacity Array: ", cap(arrStr))

	arrStr = append(arrStr, "test")
	fmt.Println("New Array: ", arrStr)

}
