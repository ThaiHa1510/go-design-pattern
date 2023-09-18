package main

import (
	"fmt"
	"sync"
)

// Step 1: Định nghĩa lớp Singleton
type Singleton struct {
	data string
}

var instance *Singleton
var once sync.Once

// Step 2: Hàm khởi tạo duy nhất của Singleton
func GetInstance(i int) *Singleton {
	once.Do(func() {
		instance = &Singleton{data: "Initial data " + fmt.Sprint(i)}

	})
	fmt.Println("Instance created i", i)
	return instance
}

// Step 3: Sử dụng Singleton
func main() {
	for i := 0; i < 100; i++ {
		singleton := GetInstance(i)
		fmt.Println(singleton.data)
	}
	fmt.Scanln()
}
