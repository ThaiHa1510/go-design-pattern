package main

import "fmt"

type House struct {
	Wall     string
	Windows  string
	Floors   int
	Bedrooms int
}

type HouseBuilder struct {
	house *House
}

func (b *HouseBuilder) BuildWall(wall string) *HouseBuilder {
	b.house.Wall = "Wall " + wall
	return b
}

func (b *HouseBuilder) BuildWindows(window string) *HouseBuilder {
	b.house.Windows = "Windows " + window
	return b
}

func (b *HouseBuilder) BuildFloors(floors int) *HouseBuilder {
	b.house.Floors = floors
	return b
}

func (b *HouseBuilder) BuildBedrooms(bedrooms int) *HouseBuilder {
	b.house.Bedrooms = bedrooms
	return b
}

func (b *HouseBuilder) Build() *House {
	return b.house
}

func NewBuilder() HouseBuilder {
	return HouseBuilder{
		house: &House{},
	}
}

// Director
type director struct {
	builder HouseBuilder
}

func NewDirector(b HouseBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) BuildHouse() *House {
	d.builder.BuildWall("Wood").
		BuildWindows("Windows").
		BuildFloors(3).
		BuildBedrooms(4)
	return d.builder.Build()
}

func main() {
	director := NewDirector(NewBuilder())
	house := director.BuildHouse()
	fmt.Println(house)
}
