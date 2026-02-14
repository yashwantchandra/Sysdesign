package main

import "fmt"

type House struct {
	Walls   string
	Doors   int
	Windows int
	Roof    string
}

type HouseBuilder struct {
	house House
}

func NewHouseBuilder() *HouseBuilder {
	return &HouseBuilder{}
}

func (b *HouseBuilder) SetWalls(w string) *HouseBuilder {
	b.house.Walls = w
	return b
}

func (b *HouseBuilder) SetDoors(n int) *HouseBuilder {
	b.house.Doors = n
	return b
}

func (b *HouseBuilder) SetWindows(n int) *HouseBuilder {
	b.house.Windows = n
	return b
}

func (b *HouseBuilder) SetRoof(r string) *HouseBuilder {
	b.house.Roof = r
	return b
}

func (b *HouseBuilder) Build() House {
	return b.house
}

func main() {
	house := NewHouseBuilder().
		SetWalls("brick").
		SetDoors(2).
		SetWindows(4).
		SetRoof("tile").
		Build()

	fmt.Printf("House: %+v\n", house)
}
