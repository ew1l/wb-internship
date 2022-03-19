package main

/*
	Строитель — это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово.
	Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.
*/

import "fmt"

type House struct {
	WindowType string
	DoorType   string
	NumFloor   int
}

type Builder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() House
}

type NormalBuilder struct {
	WindowType string
	DoorType   string
	NumFloor   int
}

func NewNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (nb *NormalBuilder) SetWindowType() {
	nb.WindowType = "Wooden window"
}

func (nb *NormalBuilder) SetDoorType() {
	nb.DoorType = "Wooden door"
}

func (nb *NormalBuilder) SetNumFloor() {
	nb.NumFloor = 2
}

func (nb *NormalBuilder) GetHouse() House {
	return House{
		WindowType: nb.WindowType,
		DoorType:   nb.DoorType,
		NumFloor:   nb.NumFloor,
	}
}

type IglooBuilder struct {
	WindowType string
	DoorType   string
	NumFloor   int
}

func NewIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (ib *IglooBuilder) SetWindowType() {
	ib.WindowType = "Snow window"
}

func (ib *IglooBuilder) SetDoorType() {
	ib.DoorType = "Snow door"
}

func (ib *IglooBuilder) SetNumFloor() {
	ib.NumFloor = 1
}

func (ib *IglooBuilder) GetHouse() House {
	return House{
		WindowType: ib.WindowType,
		DoorType:   ib.DoorType,
		NumFloor:   ib.NumFloor,
	}
}

type Director struct {
	Builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{
		Builder: b,
	}
}

func (d *Director) SetBuilder(b Builder) {
	d.Builder = b
}

func (d *Director) BuildHouse() House {
	d.Builder.SetWindowType()
	d.Builder.SetDoorType()
	d.Builder.SetNumFloor()

	return d.Builder.GetHouse()
}

func main() {
	normalBuilder := NewNormalBuilder()
	iglooBuilder := NewIglooBuilder()

	director := NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Printf("Normal House Window Type: %s\n", normalHouse.WindowType)
	fmt.Printf("Normal House Door Type: %s\n", normalHouse.DoorType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.NumFloor)

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()

	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.WindowType)
	fmt.Printf("Igloo House Door Type: %s\n", iglooHouse.DoorType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.NumFloor)
}

// Normal House Window Type: Wooden window
// Normal House Door Type: Wooden door
// Normal House Num Floor: 2
// Igloo House Window Type: Snow window
// Igloo House Door Type: Snow door
// Igloo House Num Floor: 1
