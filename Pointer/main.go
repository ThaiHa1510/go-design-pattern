package main

import "fmt"

// Originator
type Distance struct {
	value int
}

func NewDistanceWithValue(value int) *Distance {
	return &Distance{value: value}
}

func NewDistance() *Distance {
	return &Distance{value: 100}
}

func NewDistanceInitial() *Distance {
	distance := Distance{}
	return &distance
}

func (distance *Distance) GetValue() int {
	return distance.value
}

func (distance *Distance) SetValue(value int) {
	distance.value = value
}
func main() {
	// set distance
	distance := NewDistance()
	distance.SetValue(1)

	//
	distance2 := NewDistance()
	distance2.SetValue(2)
	fmt.Println("distance at 2", distance2.GetValue())
	fmt.Println("distance at 1", distance2.GetValue())

	// init object
	distance3 := NewDistanceInitial()
	distance3.SetValue(3)
	fmt.Println("distance at 3", distance3.GetValue())
}
