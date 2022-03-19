package main

import "fmt"

func main() {
	a, b := 1, 2

	// Меняем местами значения a и b
	a, b = b, a

	fmt.Printf("a = %d\nb = %d\n", a, b)
}

// a = 2
// b = 1
