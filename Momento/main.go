package main

import (
	"fmt"
)

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

func (distance *Distance) RestoreValue(DistanceMomento *DistanceMomento) {
	distanceSaved := DistanceMomento.GetDistance()
	distance.value = distanceSaved.GetValue()
}

func (distance *Distance) BackupValue(DistanceMomento *DistanceMomento) {
	DistanceMomento.SetValue(*distance)
}

func (distance *Distance) GetValue() int {
	return distance.value
}
func (distance *Distance) SetValue(value int) {
	distance.value = value
}

// Momento
type DistanceMomento struct {
	distance Distance
}

func (distanceMomento *DistanceMomento) GetDistance() *Distance {
	return &distanceMomento.distance
}

func NewDistanceMomento() *DistanceMomento {
	return &DistanceMomento{
		distance: *NewDistance(),
	}
}

func (distanceMomento *DistanceMomento) SetValue(distance Distance) {
	distanceMomento.distance = distance
}

// Caretaker
type DistanceCaretaker struct {
	momento map[string]*DistanceMomento
}

func NewDistanceCaretaker() *DistanceCaretaker {
	return &DistanceCaretaker{
		momento: make(map[string]*DistanceMomento),
	}
}

func (caretaker *DistanceCaretaker) AddMomento(key string, distance Distance) {
	momento := NewDistanceMomento()
	momento.SetValue(distance)
	caretaker.momento[key] = momento
}

func (caretaker *DistanceCaretaker) GetMomento(key string) *DistanceMomento {
	return caretaker.momento[key]
}

func main() {
	distance := NewDistance()
	distanceCaretaker := NewDistanceCaretaker()
	distance.SetValue(100)
	distanceCaretaker.AddMomento("1", *distance)
	distance.SetValue(90)
	distanceCaretaker.AddMomento("2", *distance)
	distance.SetValue(80)
	distanceCaretaker.AddMomento("3", *distance)
	distance.SetValue(70)
	distanceCaretaker.AddMomento("4", *distance)
	distance.SetValue(60)
	distanceCaretaker.AddMomento("5", *distance)
	distance.SetValue(50)
	distanceCaretaker.AddMomento("6", *distance)
	distance.SetValue(40)
	distanceCaretaker.AddMomento("7", *distance)
	//	distanceMomento := distanceCaretaker.GetMomento("1")
	distanceMomento2 := distanceCaretaker.GetMomento("2")
	distance2 := distanceMomento2.GetDistance()
	//
	distanceMomento3 := distanceCaretaker.GetMomento("3")
	distance3 := distanceMomento3.GetDistance()
	//
	distanceMomento4 := distanceCaretaker.GetMomento("4")
	distance4 := distanceMomento4.GetDistance()
	//
	distanceMomento5 := distanceCaretaker.GetMomento("5")
	distance5 := distanceMomento5.GetDistance()
	fmt.Println("distance at 2", distance2.GetValue())
	fmt.Println("distance at 3", distance3.GetValue())
	fmt.Println("distance at 4", distance4.GetValue())
	fmt.Println("distance at 5", distance5.GetValue())
}
