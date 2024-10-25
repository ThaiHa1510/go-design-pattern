package main

import (
	"fmt"
	"reflect"
)

//
// Function to find the first and last occurrences of an element in a sorted array
// Step:
//	1. Find the biggest element in the array
//  2. Find the next biggest element (reduce the array)
// ...

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Hoán đổi giá trị của arr[j] và arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func findFirstAndLast(arr []int, val int) (int, int) {
	last, first := -1, -1
	for i, v := range arr {
		if v == val {
			last = i
			if first == -1 {
				first = i
			}
		}
	}
	return last, first
}

// firstFunc is a recursive function that finds the first occurrence of a target element in a sorted array.
//
// Parameters:
// - arr: the sorted array to search in.
// - low: the starting index of the search range.
// - hight: the ending index of the search range.
// - x: the target element to search for.
// - n: the length of the array.
//
// Returns:
// - The index of the first occurrence of the target element in the array, or -1 if the element is not found.
func firstFunc(arr []int, low, hight, x, n int) int {
	if low > hight {
		return -1
	}
	mid := low + (hight-low)/2
	if (mid == 0 || arr[mid-1] < x) && arr[mid] == x {
		return mid
	}
	if arr[mid] > x {
		return firstFunc(arr, low, mid-1, x, n)
	} else {
		return firstFunc(arr, mid+1, hight, x, n)
	}
}

// lastFunc is a recursive function that finds the last occurrence of a target element in a sorted array.
//
// Parameters:
// - arr: the sorted array to search in.
// - low: the starting index of the search range.
// - hight: the ending index of the search range.
// - x: the target element to search for.
// - n: the length of the array.
//
// Returns:
// - The index of the last occurrence of the target element in the array, or -1 if the element is not found.
func lastFunc(arr []int, low, hight, x, n int) int {
	if low > hight {
		return -1
	}
	mid := low + (hight-low)/2
	if (mid == n-1 || arr[mid+1] > x) && arr[mid] == x {
		return mid
	}
	if arr[mid] > x {
		return lastFunc(arr, low, mid-1, x, n)
	} else {
		return lastFunc(arr, mid+1, hight, x, n)
	}
}

func main() {
	arr := []int{1, 2, 2, 2, 2, 3, 4, 7, 8, 8}
	x := 8
	last, first := findFirstAndLast(arr, x)
	fmt.Printf("first = %d, last = %d\n", first, last)
	first2 := firstFunc(arr, 0, len(arr)-1, x, len(arr))
	fmt.Println("Using first:", first2)
	last2 := lastFunc(arr, 0, len(arr)-1, x, len(arr))
	fmt.Println("Using last:", last2)
	// Mảng số nguyên cần sắp xếp
	numbers := []int{64, 34, 25, 12, 22, 11, 90}

	// Gọi hàm sắp xếp
	bubbleSort(numbers)

	// In ra mảng sau khi sắp xếp
	fmt.Println("Mảng sau khi sắp xếp:", numbers)
	fmt.Println(reflect.TypeOf(numbers))
}
