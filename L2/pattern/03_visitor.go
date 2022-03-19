package main

/*
	Посетитель — это поведенческий паттерн проектирования, который позволяет создавать новые операции, не меняя классы объектов, над которыми эти операции могут выполняться.
*/

import (
	"fmt"
	"math"
)

type Visitor interface {
	VisitForSquare(*Square)
	VisitForCircle(*Circle)
	VisitForRectangle(*Rectangle)
}

type Shape interface {
	Accept(Visitor)
}

type Square struct {
	Side float64
}

func NewSquare(side float64) *Square {
	return &Square{
		Side: side,
	}
}

func (s *Square) Accept(v Visitor) {
	v.VisitForSquare(s)
}

type Circle struct {
	Radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{
		Radius: radius,
	}
}

func (c *Circle) Accept(v Visitor) {
	v.VisitForCircle(c)
}

type Rectangle struct {
	Length, Width float64
}

func NewRectangle(length, width float64) *Rectangle {
	return &Rectangle{
		Length: length,
		Width:  width,
	}
}

func (r *Rectangle) Accept(v Visitor) {
	v.VisitForRectangle(r)
}

type AreaCalculator struct {
	Area float64
}

func NewAreaCalculator() *AreaCalculator {
	return &AreaCalculator{}
}

func (ac *AreaCalculator) VisitForSquare(s *Square) {
	fmt.Printf("Sqaure area: %f\n", math.Pow(s.Side, 2))
}

func (ac *AreaCalculator) VisitForCircle(c *Circle) {
	fmt.Printf("Circle area: %f\n", math.Pi*math.Pow(c.Radius, 2))
}

func (ac *AreaCalculator) VisitForRectangle(r *Rectangle) {
	fmt.Printf("Rectangle area: %f\n", r.Length*r.Width)
}

func main() {
	square := NewSquare(6.32)
	circle := NewCircle(3.24)
	rectangle := NewRectangle(9.65, 4.56)

	areaCalculator := NewAreaCalculator()

	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)
}

// Sqaure area: 39.942400
// Circle area: 32.979183
// Rectangle area: 44.004000
