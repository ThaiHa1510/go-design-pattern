package main

import "fmt"

func lastFunc1(arr []int, low, hight, x, n int) int {
	if low > hight {
		return -1
	}
	mid := low + (hight-low)/2
	if (mid == hight || arr[mid+1] > x) && arr[mid] == x {
		return mid
	}
	if arr[mid] > x {
		return lastFunc1(arr, low, mid-1, x, n)
	} else {
		return lastFunc1(arr, mid+1, hight, x, n)
	}
}
func main() {
	arr := []int{1, 2, 2, 2, 2, 3, 4, 7, 8, 8}
	x := 2
	first := lastFunc1(arr, 0, len(arr)-1, x, len(arr))
	fmt.Printf("first = %d", first)
}
