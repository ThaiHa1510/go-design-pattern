package main

import (
	"fmt"
	"math"
)

func rect(arr []int, index int) int {
	if index >= len(arr)-1 {
		return 0
	}
	rect2 := rect(arr, index+2)
	total := arr[index] + rect2
	return max(total, rect(arr, index+1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minium(arr []int, index int) int {
	min := math.MaxInt64
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func main() {
	input := []int{5, 5, 10, 100, 10, 5}
	fmt.Println(rect(input, 0))
}
