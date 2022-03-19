package main

import (
	"fmt"
	"math/big" // Пакет для работы с большими числами
)

// Суммирование
func Add(a, b int64) *big.Int {
	return new(big.Int).Add(big.NewInt(a), big.NewInt(b))
}

// Вычитание
func Sub(a, b int64) *big.Int {
	return new(big.Int).Sub(big.NewInt(a), big.NewInt(b))
}

// Умножение
func Mul(a, b int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(a), big.NewInt(b))
}

// Деление
func Div(a, b int64) *big.Int {
	return new(big.Int).Div(big.NewInt(a), big.NewInt(b))
}

func main() {
	var a, b int64 = 48e17, 14e16 // a = 4800000000000000000, b = 140000000000000000

	fmt.Printf("%d + %d = %d\n", a, b, Add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, Sub(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, Mul(a, b))
	fmt.Printf("%d / %d = %d\n", a, b, Div(a, b))
}

// 4800000000000000000 + 140000000000000000 = 4940000000000000000
// 4800000000000000000 - 140000000000000000 = 4660000000000000000
// 4800000000000000000 * 140000000000000000 = 672000000000000000000000000000000000
// 4800000000000000000 / 140000000000000000 = 34
