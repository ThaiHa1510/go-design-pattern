package main

import (
	"fmt"
	"sync"
	"time"
)

type S struct {
	X int `json:"key_1"`
	Y int `json:"key_2"`
	Z int `json:"key_3"`
	W int `json:"-"`
}

func main() {
	// receive channel
	pingCh := make(chan string)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(2)
		go ping(pingCh, i, &wg)
		go pong(pingCh, i, &wg)
	}
	wg.Wait()
	var t *time.Time
	fmt.Println(t)
	var aA byte = 255
	fmt.Println(aA + 1)
	i := 42
	p := &i
	*p = 33
	fmt.Println(i)
	ch := make(chan string, 2)
	ch <- "aaa"
	fmt.Println(<-ch)
	// fmt.Println("is nil", isNil(t))
	// var a = []int{1, 2, 3, 4, 5}
	// b := a[:3]
	// b[1] = 10
	// fmt.Println(a)
	// s := S{1, 2, 3, 4}
	// b, err := json.Marshal(s)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("%s\n", s)
}

func ping(ping <-chan string, i int, wg *sync.WaitGroup) {
	message := <-ping
	fmt.Println(message)
	defer wg.Done()
}

func pong(pong chan<- string, i int, wg *sync.WaitGroup) {
	pong <- fmt.Sprintf("ping %d", i)
	defer wg.Done()
}

func isNil(val interface{}) bool {
	return val == nil
}
