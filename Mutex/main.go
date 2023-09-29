package main

import (
	"fmt"
	"sync"
)

func main() {
	// wg := sync.WaitGroup{}
	// mutex := sync.Mutex{}
	// done := make(chan bool)
	// a := 1
	// for i := 0; i < 400; i++ {
	// 	fmt.Printf("i = %d\n", i)
	// 	wg.Add(1)
	// 	go func(n int) {
	// 		fmt.Printf("n = %d\n", n)
	// 		defer wg.Done()
	// 		Add(&a, &mutex)
	// 	}(i)
	// }
	// wg.Wait()
	// fmt.Println("a =", a)
	// go func() {
	// 	time.Sleep(500 * time.Millisecond)
	// 	done <- true
	// }()

	// for {
	// 	select {
	// 	case <-done:
	// 		fmt.Println("Operation completed.")
	// 		return
	// 	case <-time.After(time.Hour):
	// 		fmt.Println("Still waiting...")
	// 	}
	// }
	// strData := "hello"
	// body := strings.NewReader(strData)
	// //or
	// //body = bytes.NewBufferString(strData)
	// resp, err := http.Post("http://localhost:8080", "text/plain", body)
	// if err != nil {
	// 	panic(err)
	// }
	//panicCode(nil)
	defer handlePanic()
	panic("This is a panic!")
	//defer resp.Body.Close()

}

func Add(a *int, mutex *sync.Mutex) {
	mutex.Lock()
	fmt.Println("Add Function")
	*a = *a + 1
	defer mutex.Unlock()
}

func panicCode(myVal interface{}) {
	a := myVal.(string)

	fmt.Println(a)
	panic("This is a panic!")
}

// BETTER
func handlePanic() {
	fmt.Println("handlePanic")
	if panicInfo := recover(); panicInfo != nil {
		fmt.Println(panicInfo)
	}
}
