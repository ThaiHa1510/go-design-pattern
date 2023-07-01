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

type TrafficCounter struct {
	value     int
	observers []Observer
}

func (trafficCounter *TrafficCounter) Notify() {
	for i := 0; i < len(trafficCounter.observers); i++ {
		trafficCounter.observers[i].Update(trafficCounter.value)
	}

}

func (trafficCounter *TrafficCounter) Attach(observer Observer) {
	trafficCounter.observers = append(trafficCounter.observers, observer)
}

func (trafficCounter *TrafficCounter) Detach(observer Observer) {
	for i := 0; i < len(trafficCounter.observers); i++ {
		if trafficCounter.observers[i] == observer {
			trafficCounter.observers = append(trafficCounter.observers[:i], trafficCounter.observers[i+1:]...)
		}
	}
}

func (trafficCounter *TrafficCounter) Update(value int) {
	trafficCounter.value = value
	trafficCounter.Notify()
}

// Observer
type SaveLog struct {
	trafficCount int
}

func (saveLog *SaveLog) Update(value int) {
	saveLog.trafficCount = value
	fmt.Printf("SaveLog: %d\n", saveLog.trafficCount)
}

func NewSaveLog() *SaveLog {
	return &SaveLog{}
}
func main() {
	fmt.Println("This is example of Observer Pattern")
	// Khơi tạo subject
	subject := &TrafficCounter{
		value:     0,
		observers: make([]Observer, 0),
	}

	// Khơi tạo observer
	observer1 := NewSaveLog()
	observer2 := NewSaveLog()
	observer3 := NewSaveLog()
	// attach subject
	subject.Attach(observer1)
	subject.Attach(observer2)
	subject.Attach(observer3)
	subject.Update(100)
	// update subject
	subject.Update(90)
	// detach observer
	subject.Detach(observer1)
	subject.Update(80)
}
