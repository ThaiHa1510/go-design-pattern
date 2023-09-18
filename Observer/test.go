package main

import (
	"fmt"
)

type Observer interface {
	Update(value int)
}

type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify()
}

type RouterCounter struct {
	value     int
	observers []Observer
}

func (routerCounter *RouterCounter) Notify() {
	fmt.Println("RouterCounter: ", len(routerCounter.observers))
	for i := 0; i < len(routerCounter.observers); i++ {
		routerCounter.observers[i].Update(routerCounter.value)
	}
}

func (routerCounter *RouterCounter) Attach(observer Observer) {
	routerCounter.observers = append(routerCounter.observers, observer)
}

func (routerCounter *RouterCounter) Detach(observer Observer) {
	// for i := 0; i < len(routerCounter.observers); i++ {
	// 	fmt.Println("RouterCounter: ", i)
	// 	if routerCounter.observers[i] == observer {
	// 		routerCounter.observers = append(routerCounter.observers[:i], routerCounter.observers[i+1:]...)
	// 	}
	// }
	for i := 0; i < len(routerCounter.observers); i++ {
		fmt.Println("RouterCounter: ", i)
		if routerCounter.observers[i] == observer {
			routerCounter.observers = append(routerCounter.observers[:i], routerCounter.observers[i+1:]...)
		}
	}
}

func (routerCounter *RouterCounter) Update(value int) {
	routerCounter.value = value
	routerCounter.Notify()
}

type SaveLog struct {
	value int
}

func (s *SaveLog) Update(value int) {
	s.value = value
	fmt.Printf("SaveLog: %d\n", value)
}

func NewLog() *SaveLog {
	return &SaveLog{}
}

func main() {
	fmt.Println("This is example of Observer Pattern")
	observer1 := NewLog()
	observer2 := NewLog()
	observer3 := NewLog()
	routerCounter := RouterCounter{}
	routerCounter.Attach(observer1)
	routerCounter.Attach(observer2)
	routerCounter.Attach(observer3)
	routerCounter.Update(100)
	/// Detach observer
	routerCounter.Detach(observer1)
	routerCounter.Update(90)
}
