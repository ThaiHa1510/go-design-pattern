package main

import (
	"encoding/json"
	"fmt"
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
func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
		fmt.Printf("Running goroutine \n")
	}()
	<-ch
	fmt.Println("End main")
}
