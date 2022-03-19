package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

// Конструктор структуры Point
func NewPoint(X, Y int) Point {
	return Point{
		x: X,
		y: Y,
	}
}

// Функция нахождения расстояния между точками
func Distance(a *Point, b *Point) float64 {
	return math.Sqrt(math.Pow(float64(b.x-a.x), 2) + math.Pow(float64(b.y-a.y), 2))
}

func main() {
	// Создаем две точки
	a := NewPoint(2, 4)
	b := NewPoint(4, 6)

	// Находим расстояние
	fmt.Printf("Расстояние между точками %v и %v: %g\n", a, b, Distance(&a, &b))
}

// Расстояние между точками {2 4} и {4 6}: 2.8284271247461903
