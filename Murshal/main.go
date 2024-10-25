package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type S struct {
	X int `json:"key_1"`
	Y int `json:"key_2",omi`
	Z int `json:"key_3"`
	W int `json:"-"`
}

func testMurshal() {
	s := S{X: 1, Y: 2, Z: 3, W: 4}
	js, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", js)
}

func maxSum(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	firstMax := arr[0]
	secondMax := 0
	for i := 1; i < len(arr); i++ {
		if firstMax > arr[i] && arr[i] > secondMax {
			secondMax = arr[i]
		}

	}
	return firstMax + secondMax
}
func main() {
	// ch := make(chan int)
	// go func() {
	// 	ch <- 1
	// 	fmt.Printf("Running goroutine \n")
	// }()
	// <-ch
	// fmt.Println("End main")
	arr := []int{5, 1, 1, 5}
	max := maxSum(arr)
	fmt.Println("value", max)
	ch1 := make(chan int, 1)
	fmt.Println("len", len(ch1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
